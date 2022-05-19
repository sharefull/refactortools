package requiredcheck

import (
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const doc = "requiredcheck is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "requiredcheck",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			decl, _ := decl.(*ast.GenDecl)
			if decl == nil || decl.Tok != token.TYPE {
				continue
			}
			for _, spec := range decl.Specs {
				spec, _ := spec.(*ast.TypeSpec)
				if spec == nil || !strings.HasSuffix(spec.Name.Name, "Request") {
					continue
				}
				typ, _ := spec.Type.(*ast.StructType)
				if typ == nil || typ.Fields == nil {
					continue
				}
				for _, field := range typ.Fields.List {
					if field.Tag == nil || field.Tag.Kind != token.STRING {
						continue
					}
					strTag, err := strconv.Unquote(field.Tag.Value)
					if err != nil {
						continue
					}
					tag := reflect.StructTag(strTag)
					binding := strings.Split(tag.Get("binding"), ",")
					var hasRequired bool
					for _, b := range binding {
						if b == "required" {
							hasRequired = true
							break
						}
					}
					if !hasRequired {
						pass.Reportf(field.Pos(), "field %s is not required", field.Names[0].Name)
					}
				}
			}
		}
	}

	return nil, nil
}
