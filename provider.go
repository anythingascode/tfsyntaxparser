package tfsyntax

import (
	"log"

	"github.com/hashicorp/hcl/v2/gohcl"
)

// Provider implement azure provider block
type ProviderConfig struct {
	// Provider name
	Name string
	// Provider alias
	Alias string
}

// Method to parse provider block
func (c *Config) GetProviderConfig(filepath string) []*ProviderConfig {
	for _, block := range c.BlocksOfType(filepath, "provider", bodySchema) {
		featuresBlock := c.BlocksWithinBlock(block, "features", providerConfigSchema)
		providerBlock := &ProviderConfig{
			Name: block.Labels[0],
		}
		if providerBlock.Name == "azurerm" && len(featuresBlock) == 0 {
			log.Fatal("No features block found!")
		}
		attr, _ := block.Body.JustAttributes()
		if _, ok := attr["alias"]; ok {
			gohcl.DecodeExpression(attr["alias"].Expr, nil, &providerBlock.Alias)
		}
		c.Providers = append(c.Providers, providerBlock)
	}
	return c.Providers
}
