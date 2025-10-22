package bigquery

import (
	_ "embed"
	"testing"
)

//go:embed testdata/script.js
var script string

func TestIntegration(t *testing.T) { //nolint:paralleltest
	//sqltest.RunScript(t, "bigquery", "bigquery://projectid/location/dataset", script)
	t.Skip()
}
