package main

import (
	"github.com/sharefull/refactortools/errisas"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(errisas.Analyzer) }
