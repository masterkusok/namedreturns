// Package analyzer contains all the analysis logic.
package analyzer

import (
	"go/ast"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:       "errcheck",
	Doc:        "check for unchecked errors",
	Run:        runAnalysis,
	ResultType: reflect.TypeOf(Result{}),
}

var (
	// allowAnonymous determines if anonymous functions must be checked.
	skipAnonymous bool

	// skipTests determines if test files must be checked.
	skipTests bool
)

func init() {
	Analyzer.Flags.BoolVar(
		&skipAnonymous,
		"a",
		false,
		"skip anonymous functions",
	)

	Analyzer.Flags.BoolVar(
		&skipTests,
		"t",
		false,
		"skip test files",
	)
}

// UnnamedReturn represents the function result token without name.
type UnnamedReturn struct {
	Line        int
	FileName    string
	Type        string
	FuncName    string
	IsAnonymous bool
}

// Result represents the result of namedreturns check run.
type Result struct {
	UnnamedReturns []UnnamedReturn
}

// runAnalysis contains main analysis pass logic.
func runAnalysis(pass *analysis.Pass) (result any, err error) {
	var unnamedReturns []UnnamedReturn
	for _, f := range pass.Files {
		file := pass.Fset.File(f.Pos())
		if isIgnoredFile(file.Name()) {
			continue
		}

		ast.Inspect(f, func(node ast.Node) (ok bool) {
			switch decl := node.(type) {
			case *ast.FuncDecl:
				unnamedReturns = append(unnamedReturns, checkFuncDecl(pass, decl)...)
			case *ast.FuncLit:
				if skipAnonymous {
					return true
				}

				unnamedReturns = append(unnamedReturns, checkFuncLit(pass, decl)...)
			}

			return true
		})
	}

	return Result{
		UnnamedReturns: unnamedReturns,
	}, nil
}

// checkFuncDecl checks function declaration node for unnamed returns.  pass and
// funcDecl must not be nil.
func checkFuncDecl(pass *analysis.Pass, funcDecl *ast.FuncDecl) (returns []UnnamedReturn) {
	if funcDecl.Type.Results == nil {
		return nil
	}

	for _, res := range funcDecl.Type.Results.List {
		if len(res.Names) != 0 {
			continue
		}

		pos := pass.Fset.Position(res.Pos())
		newRet := UnnamedReturn{
			Line:     pos.Line,
			FileName: pos.Filename,
			FuncName: funcDecl.Name.Name,
		}

		returns = append(returns, newRet)
		pass.Reportf(res.Pos(), "function %s has unnamed returns", newRet.FuncName)

		break
	}

	return returns
}

// checkFuncDecl checks ananymoys function declaration for unnamed returns.
// pass and funcDecl must not be nil.
func checkFuncLit(pass *analysis.Pass, funcDecl *ast.FuncLit) (returns []UnnamedReturn) {
	if funcDecl.Type.Results == nil {
		return nil
	}

	for _, res := range funcDecl.Type.Results.List {
		if len(res.Names) != 0 {
			continue
		}

		pos := pass.Fset.Position(res.Pos())
		newRet := UnnamedReturn{
			Line:        pos.Line,
			FileName:    pos.Filename,
			IsAnonymous: true,
		}

		returns = append(returns, newRet)
		pass.Reportf(res.Pos(), "anonymous function has unnamed returns")

		break
	}

	return returns
}

// isIgnoredFile returns true if file should be ignored by check.  Such kind of
// files can be code generated from protobuf.
func isIgnoredFile(filename string) (ok bool) {
	// TODO(f.setrakov): Refactor.
	ignoredSuffixes := []string{"gen.go", "pb.go"}
	if skipTests {
		ignoredSuffixes = append(ignoredSuffixes, "test.go")
	}

	for _, suffix := range ignoredSuffixes {
		if strings.HasSuffix(filename, suffix) {
			return true
		}
	}

	return false
}
