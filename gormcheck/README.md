# sharefull/refactortools/gormcheck

[![pkg.go.dev][gopkg-badge]][gopkg]

`gormcheck` finds code which may be SQL injection.

## Install

You can get `gormcheck` by `go install` command (Go 1.16 and higher).

```bash
$ go install github.com/sharefull/refactortools/gormcheck/cmd/gormcheck@latest
```

## How to use

`gormcheck` run with `go vet` as below when Go is 1.12 and higher.

```bash
$ go vet -vettool=$(which gormcheck) ./...
```

## Analyze with golang.org/x/tools/go/analysis

You can use [gormcheck.Analyzers](https://pkg.go.dev/github.com/sharefull/refactortools/gormcheck/#Analyzers) with [unitchecker](https://golang.org/x/tools/go/analysis/unitchecker).

```go
package main

import (
	"github.com/sharefull/refactortools/gormcheck"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(gormcheck.Analyzers...) }
```

The gormcheck provides following analyzers.

* [gorminjection.Analyzer](https://pkg.go.dev/github.com/sharefull/refactortools/gormcheck/passes/gorminjection#Analyzer)

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/sharefull/refactortools/gormcheck
[gopkg-badge]: https://pkg.go.dev/badge/github.com/sharefull/refactortools/gormcheck?status.svg
