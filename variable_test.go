package tfsyntax

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_BlocksOfType(t *testing.T) {
	var c Config
	blocks := c.BlocksOfType("testdata/main.tf", "variable", bodySchema)
	if len(blocks) == 0 {
		t.Errorf("list iof blocks wanted , got an empty list")
	}
}

func Test_GetVariablesConfig(t *testing.T) {
	var c Config
	var gotNames []string
	var gotTypes []string
	var gotDescriptions []string
	gotVariables := c.GetVariablesConfig("testdata/main.tf")
	for _, g := range gotVariables {
		gotNames = append(gotNames, g.Name)
		gotTypes = append(gotTypes, g.Type)
		gotDescriptions = append(gotDescriptions, g.Description)
		//gotValidation = append(gotValidation, g.Validation.Condition)
	}
	var wantNames = []string{"location", "test", "validation"}
	if !cmp.Equal(wantNames, gotNames) {
		t.Errorf("got %v, wanted %v", gotNames, wantNames)
	}
	var wantTypes = []string{"string", "string", ""}
	if !cmp.Equal(wantTypes, gotTypes) {
		t.Errorf("got %v, wanted %v", gotTypes, wantTypes)
	}
	var wantDescriptions = []string{"(Required) Specifies the supported Azure location for Diagnostic Setting module.", "", ""}
	if !cmp.Equal(wantDescriptions, gotDescriptions) {
		t.Errorf("got %v, wanted %v", gotDescriptions, wantDescriptions)
	}

}
