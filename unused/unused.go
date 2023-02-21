package unused

import (
	"go/token"
	"go/types"
	"log"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"github.com/sharefull/refactortools/unused/internal"
	"golang.org/x/exp/maps"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/objectpath"
)

var (
	Uses       = make(map[string]types.Object)
	Defs       = make(map[string]types.Object)
	Interfaces map[string]*types.Interface
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
	for _, obj := range pass.TypesInfo.Uses {
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

		if strings.Contains(file.Name(), "moqs") {
			continue
		}

		if obj.Exported() {
			id := obj.Pkg().Path() + "." + obj.Name()
			Uses[id] = obj
		}
	}

	for _, obj := range pass.TypesInfo.Defs {
		if obj == nil {
			continue
		}

		objVar, _ := obj.(*types.Var)
		if objVar != nil && objVar.IsField() {
			continue
		}

		file := pass.Fset.File(obj.Pos())
		if strings.HasSuffix(file.Name(), "_test.go") {
			continue
		}

		if strings.Contains(file.Name(), "moqs") {
			continue
		}

		if obj.Exported() {
			id := obj.Pkg().Path() + "." + obj.Name()
			p, _ := objectpath.For(obj)
			log.Println("p:", p, "id:", id)
			Defs[id] = obj
		}
	}

	Interfaces = make(map[string]*types.Interface)
	maps.Copy(Interfaces, analysisutil.Interfaces(pass.Types))

	return nil
}
