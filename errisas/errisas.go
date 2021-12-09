package errisas

import (
	"go/token"
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const doc = "errisas finds error handling codes which do not use errors.Is or errors.As"

var Analyzer = &analysis.Analyzer{
	Name: "errisas",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

var (
	errType     = types.Universe.Lookup("error").Type()
	errIterface = errType.Underlying().(*types.Interface)
)

func run(pass *analysis.Pass) (interface{}, error) {
	checkIs(pass)
	checkAs(pass)
	return nil, nil
}

func checkIs(pass *analysis.Pass) {
	s := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	analysisutil.InspectFuncs(s.SrcFuncs, func(_ int, instr ssa.Instruction) bool {
		binop, _ := instr.(*ssa.BinOp)
		if binop == nil {
			return true
		}
		if binOpErrNil(binop) {
			pass.Reportf(binop.Pos(), "must use errors.Is")
		}
		return true
	})
}

func binOpErrNil(binop *ssa.BinOp) bool {

	if binop.Op != token.EQL &&
		binop.Op != token.NEQ {
		return false
	}

	if !types.Identical(binop.X.Type(), errType) &&
		!types.Identical(binop.Y.Type(), errType) {
		return false
	}

	xIsConst, yIsConst := isConst(binop.X), isConst(binop.Y)
	switch {
	case !xIsConst && yIsConst: // err != nil or err == nil
		return false
	case xIsConst && !yIsConst: // nil != err or nil == err
		return false
	}

	return true
}

func isConst(v ssa.Value) bool {
	_, ok := v.(*ssa.Const)
	return ok
}

func checkAs(pass *analysis.Pass) {
	s := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	analysisutil.InspectFuncs(s.SrcFuncs, func(_ int, instr ssa.Instruction) bool {
		if isTarget(instr) {
			pass.Reportf(instr.Pos(), "must use errors.As")
		}
		return true
	})
}

func isTarget(instr ssa.Instruction) bool {
	typassert, _ := instr.(*ssa.TypeAssert)
	if typassert == nil {
		return false
	}

	return types.Identical(typassert.X.Type(), errType)
}
