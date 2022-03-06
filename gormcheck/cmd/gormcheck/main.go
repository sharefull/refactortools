package main

import (
	"github.com/sharefull/refactortools/gormcheck"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(gormcheck.Analyzers...) }
