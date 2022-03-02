# sharefull/refactortools/gormcheck

[![pkg.go.dev][gopkg-badge]][gopkg]

`gormcheck` finds code which may be SQL injection.

## Install

You can get `errisas` by `go install` command (Go 1.16 and higher).

```bash
$ go install github.com/sharefull/refactortools/gormcheck/cmd/gormcheck@latest
```

## How to use

`gormcheck` run with `go vet` as below when Go is 1.12 and higher.

```bash
$ go vet -vettool=$(which gormcheck) ./...
```

## Analyze with golang.org/x/tools/go/analysis

You can use [gormcheck.Analyzer](https://pkg.go.dev/github.com/sharefull/refactortools/gormcheck/#Analyzer) with [unitchecker](https://golang.org/x/tools/go/analysis/unitchecker).

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/sharefull/refactortools/gormcheck
[gopkg-badge]: https://pkg.go.dev/badge/github.com/sharefull/refactortools/gormcheck?status.svg
