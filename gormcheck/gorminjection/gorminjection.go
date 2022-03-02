package gorminjection

import (
	"fmt"
	"go/ast"
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "gorminjection is ..."

var Analyzer = &analysis.Analyzer{
	Name: "gorminjection",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var gormPaths = []string{
	"gorm.io/gorm",
	"github.com/jinzhu/gorm",
}

var targetFuncs = []string{"Find", "First", "Delete", "Last", "Take"}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	var dbs []types.Object
	for _, path := range gormPaths {
		db := analysisutil.LookupFromImports(pass.Pkg.Imports(), path, "DB")
		if db != nil {
			dbs = append(dbs, db)
		}
	}

	// skip (does not use Gorm)
	if len(dbs) == 0 {
		return nil, nil
	}

	funcs := make([]*types.Func, 0, len(dbs)*len(targetFuncs))
	for _, db := range dbs {
		// type of *gorm.DB
		dbPtr := types.NewPointer(db.Type())

		// methods of *gorm.DB
		for _, fname := range targetFuncs {
			f := analysisutil.MethodOf(dbPtr, fname)
			if f == nil {
				return nil, fmt.Errorf("cannot find (*gorm.DB).%s", fname)
			}

			funcs = append(funcs, f)
		}
	}

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call, _ := n.(*ast.CallExpr)
		if call == nil {
			return
		}

		fun := getFun(pass, call.Fun)
		for _, f := range funcs {
			if fun != f || len(call.Args) != 2 {
				continue
			}

			typArg2 := pass.TypesInfo.TypeOf(call.Args[1])
			if types.Identical(typArg2, types.Typ[types.String]) {
				pass.Reportf(call.Pos(), "it may be SQL injection")
				break
			}
		}
	})

	return nil, nil
}

func getFun(pass *analysis.Pass, fun ast.Expr) *types.Func {
	switch fun := fun.(type) {
	case *ast.Ident:
		obj, _ := pass.TypesInfo.ObjectOf(fun).(*types.Func)
		return obj
	case *ast.SelectorExpr:
		obj, _ := pass.TypesInfo.ObjectOf(fun.Sel).(*types.Func)
		return obj
	}
	return nil
}
