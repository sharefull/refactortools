# sharefull/refactortools/errisas

[![pkg.go.dev][gopkg-badge]][gopkg]

`errisas` finds error handling codes which do not use errors.Is or errors.As.

## Install

You can get `errisas` by `go install` command (Go 1.16 and higher).

```bash
$ go install github.com/sharefull/refactortools/errisas/cmd/errisas@latest
```

## How to use

`dive` run with `go vet` as below when Go is 1.12 and higher.

```bash
$ go vet -vettool=$(which errisas) ./...
```

## Analyze with golang.org/x/tools/go/analysis

You can use [errisas.Analyzer](https://pkg.go.dev/github.com/sharefull/refactortools/errisas/#Analyzer) with [unitchecker](https://golang.org/x/tools/go/analysis/unitchecker).

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/sharefull/refactortools/errisas
[gopkg-badge]: https://pkg.go.dev/badge/github.com/sharefull/refactortools/errisas?status.svg
