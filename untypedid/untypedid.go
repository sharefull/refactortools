package untypedid

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "untypedid find ids whose type is not user defined type"

var Analyzer = &analysis.Analyzer{
	Name: "untypedid",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		id, _ := n.(*ast.Ident)
		if id == nil || (!strings.HasSuffix(id.Name, "ID") &&
			!strings.HasSuffix(id.Name, "Id")) {
			return
		}

		typ, _ := pass.TypesInfo.TypeOf(id).(*types.Basic)
		if typ == nil {
			return
		}

		pass.Reportf(n.Pos(), "%s's type should be user defined type but %s", id.Name, typ)
	})

	return nil, nil
}
