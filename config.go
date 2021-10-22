package tfsyntax

import (
	"log"
	"os"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// Config type implement all terraform blocks
type Config struct {
	Variables []*Variable
	Modules   []*ModuleConfig
	Providers []*ProviderConfig
}

// Method to read HCL file
func (c *Config) HclFile(filePath string) *hcl.File {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("read error ", filePath)
	}
	file, diags := hclsyntax.ParseConfig(content, filePath, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		log.Fatal("parse config error", diags)
	}
	return file
}

// Method to get HCL block of given type
func (c *Config) BlocksOfType(filePath, typeName string, BodySchema *hcl.BodySchema) hcl.Blocks {
	file := c.HclFile(filePath)
	bc, _, diags := file.Body.PartialContent(BodySchema)
	if diags.HasErrors() {
		log.Fatal("file content error ", diags)
	}
	return bc.Blocks.OfType(typeName)
}

// Method to get blocks within a block
func (c *Config) BlocksWithinBlock(block *hcl.Block, typeName string, BodySchema *hcl.BodySchema) hcl.Blocks {
	bc, _, _ := block.Body.PartialContent(BodySchema)
	return bc.Blocks.OfType(typeName)
}

// HCL write implementation
func (c *Config) HclWriteFile(filePath string) *hclwrite.File {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("read error ", filePath)
	}
	file, diags := hclwrite.ParseConfig(content, filePath, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		log.Fatal("parse config error", diags)
	}
	return file
}

// HCL write implementation
func (c *Config) HclWriteBlockOfType(filePath, typeName string) []*hclwrite.Block {
	var BlocksByType []*hclwrite.Block
	file := c.HclWriteFile(filePath)
	for _, block := range file.Body().Blocks() {
		if block.Type() == typeName {
			BlocksByType = append(BlocksByType, block)
		}
	}
	return BlocksByType
}
