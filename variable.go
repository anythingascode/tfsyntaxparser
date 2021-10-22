package tfsyntax

import (
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

// Variable implement variable block
type Variable struct {
	// Variable name
	Name string
	// Variable description
	Description string
	// Variable default value
	Default string
	// Variable Type
	Type string
	// Variable sensitive bool
	Sensitive bool
	// Variable validaion block, It is not not part of validation.
	// added for future requirements
	Validation *VariableValidation
}

//added for future requirements
type VariableValidation struct {
	Condition    hcl.Expression
	ErrorMessage string
}

// Method to parse variable block
func (c *Config) GetVariablesConfig(filePath string) []*Variable {
	for _, block := range c.BlocksOfType(filePath, "variable", bodySchema) {
		variable := &Variable{
			Name: block.Labels[0],
		}
		attr, _ := block.Body.JustAttributes()
		if _, ok := attr["description"]; ok {
			gohcl.DecodeExpression(attr["description"].Expr, nil, &variable.Description)
		}
		if _, ok := attr["sensitive"]; ok {
			gohcl.DecodeExpression(attr["sensitive"].Expr, nil, &variable.Sensitive)
		}
		if _, ok := attr["default"]; ok {
			gohcl.DecodeExpression(attr["default"].Expr, nil, &variable.Default)
		}
		if _, ok := attr["type"]; ok {
			variable.Type = hcl.ExprAsKeyword(attr["type"].Expr)
		}
		c.Variables = append(c.Variables, variable)
	}
	return c.Variables
}

// Method to parse validation block
func (c *Config) GetVariableValidationConfig(filePath, varName string) hcl.Expression {
	var validation = &VariableValidation{}
	for _, block := range c.BlocksOfType(filePath, "variable", bodySchema) {
		validationBlocks := c.BlocksWithinBlock(block, "validation", variableValidationBlockSchema)
		if block.Labels[0] == varName && len(validationBlocks) > 0 {
			for _, vb := range validationBlocks {
				attr, _ := vb.Body.JustAttributes()
				if _, ok := attr["condition"]; ok {
					validation.Condition = attr["condition"].Expr
				}
				if _, ok := attr["error_message"]; ok {
					gohcl.DecodeExpression(attr["error_message"].Expr, nil, &validation.ErrorMessage)
				}
			}
		}
	}
	return validation.Condition
}
