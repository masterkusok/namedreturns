## namedreturns

namedreturns is a static check that finds all returns without names in Go code.

All returns must be named:
``` go
// both of return parameters do not have names.
func foo() (int, error) {
    ...
}

// must be
func foo() (bar int, err error) {
    ...
}
```

## Install

``` sh
go install github.com/masterkusok/namedreturns@latest
```

## Usage

``` sh
namedreturns [flags] [files]
```

### Flags

List of available flags now:
- [a](#anonymous): disables check for anonymous functions.
- [t](#test): disables check for test files.

#### <a href="#anonymous" id="anonymous" name="anonymous">`a (anonymous)`</a>

Flag a disables check for anonymous functions (callbacks):
``` go
func foo() {
    // callback is an anonymous function, flag -a disables check for such kind of declarations
    callback := func() (int, error) {
        ...
    }
    ...
}
```

#### <a href="#test" id="test" name="test">`test (test)`</a>

Flag a disables check for function in test files.  Test files are identified by
`test.go` suffix.