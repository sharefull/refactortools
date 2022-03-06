package gormcheck

import (
	"github.com/sharefull/refactortools/gormcheck/passes/gorminjection"
	"golang.org/x/tools/go/analysis"
)

// Analyzers is a list of analyzers in gormcheck.
// It can be use as an argument of unitchecker.Main.
//
//   package main
//
//   import (
//   	"github.com/sharefull/refactortools/gormcheck"
//   	"golang.org/x/tools/go/analysis/unitchecker"
//   )
//
//   func main() { unitchecker.Main(gormcheck.Analyzers...) }
//
var Analyzers = []*analysis.Analyzer{
	gorminjection.Analyzer,
}
