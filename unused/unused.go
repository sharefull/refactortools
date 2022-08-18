package unused

import (
	"fmt"
	"go/token"
	"go/types"
	"strings"

	"github.com/sharefull/refactortools/unused/internal"
	"golang.org/x/tools/go/packages"
)

var (
	Uses = make(map[string]types.Object)
	Defs = make(map[string]types.Object)
)

var Analyzer = &internal.Analyzer{
	Name: "unused",
	Doc:  "unused is ...",
	Config: &packages.Config{
		Fset: token.NewFileSet(),
		Mode: packages.NeedName | packages.NeedTypes |
			packages.NeedSyntax | packages.NeedTypesInfo |
			packages.NeedModule | packages.NeedDeps | packages.NeedImports,
	},
	Run: run,
}

func run(pass *internal.Pass) error {
	fmt.Println(pass.PkgPath)
	fmt.Println("Uses")
	for id, obj := range pass.TypesInfo.Uses {
		fmt.Println("\t", id)
		if obj == nil {
			continue
		}

		file := pass.Fset.File(obj.Pos())
		if file == nil {
			continue
		}

		if strings.HasSuffix(file.Name(), "_test.go") {
			continue
		}

		if obj.Exported() {
			id := obj.Pkg().Path() + "." + obj.Name()
			Uses[id] = obj
		}
	}

	fmt.Println("Defs")
	for id, obj := range pass.TypesInfo.Defs {
		fmt.Println("\t", id)
		if obj == nil {
			continue
		}

		file := pass.Fset.File(obj.Pos())
		if strings.HasSuffix(file.Name(), "_test.go") {
			continue
		}

		if obj.Exported() {
			id := obj.Pkg().Path() + "." + obj.Name()
			Defs[id] = obj
		}
	}

	return nil
}