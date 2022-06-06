package main

import (
	"untypedid"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(untypedid.Analyzer) }
