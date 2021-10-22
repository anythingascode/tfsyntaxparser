package tfsyntax

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_GetProviderConfig(t *testing.T) {
	var c Config
	var gotProviders = []ProviderConfig{}
	for _, v := range c.GetProviderConfig("testdata/main.tf") {
		var gotProvider = ProviderConfig{
			Name:  v.Name,
			Alias: v.Alias,
		}
		gotProviders = append(gotProviders, gotProvider)
	}
	var wantProviders = []ProviderConfig{
		{
			Name:  "hashicups",
			Alias: "test",
		},
		{
			Name:  "azurerm",
			Alias: "testazure",
		},
	}
	if !cmp.Equal(gotProviders, wantProviders) {
		t.Errorf("got %v, wanted %v", gotProviders, wantProviders)
	}
}
