package main

import (
	"github.com/sharefull/refactortools/gormcheck/gorminjection"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(gorminjection.Analyzer) }
