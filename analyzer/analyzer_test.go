package analyzer_test

import (
	"testing"

	"github.com/masterkusok/namedreturns/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestValid(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.Analyzer, "valid")
}

func TestInvalid(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.Analyzer, "invalid")
}

func TestAnonymous(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer.Analyzer, "anonymous")
}
