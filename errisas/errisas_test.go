package errisas_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/sharefull/refactortools/errisas"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, errisas.Analyzer, "a")
}
