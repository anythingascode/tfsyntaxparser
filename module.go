package tfsyntax

import "github.com/hashicorp/hcl/v2/gohcl"

// ModuleConfig represents a "module" block.
type ModuleConfig struct {
	Name    string
	Source  string
	Version string
}

// Method to parse module block
func (c *Config) GetModuleConfig(filepath string) []*ModuleConfig {
	for _, block := range c.BlocksOfType(filepath, "module", bodySchema) {
		moduleBlock := &ModuleConfig{
			Name: block.Labels[0],
		}
		attr, _ := block.Body.JustAttributes()
		if _, ok := attr["source"]; ok {
			gohcl.DecodeExpression(attr["source"].Expr, nil, &moduleBlock.Source)
		}
		if _, ok := attr["version"]; ok {
			gohcl.DecodeExpression(attr["version"].Expr, nil, &moduleBlock.Version)
		}
		c.Modules = append(c.Modules, moduleBlock)
	}
	return c.Modules
}
