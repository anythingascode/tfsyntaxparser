package tfsyntax

import (
	//"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_GetModuleConfig(t *testing.T) {
	var c Config
	var gotModules = []ModuleConfig{}
	for _, v := range c.GetModuleConfig("testdata/main.tf") {
		var gotModule = ModuleConfig{
			Name:    v.Name,
			Source:  v.Source,
			Version: v.Version,
		}
		gotModules = append(gotModules, gotModule)
	}
	var wantModules = []ModuleConfig{
		{
			Name:    "psl1",
			Source:  "./coffee1",
			Version: "1.0.1",
		},
		{
			Name:    "psl2",
			Source:  "./coffee2",
			Version: "1.0.2",
		},
		{
			Name:    "psl3",
			Source:  "./coffee3",
			Version: "1.0.3",
		},
	}
	if !cmp.Equal(gotModules, wantModules) {
		t.Errorf("got %v, wanted %v", gotModules, wantModules)
	}
}
