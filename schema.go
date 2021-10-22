package tfsyntax

import (
	hcl "github.com/hashicorp/hcl/v2"
)

// Terraform HCL schema
var bodySchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "terraform",
			LabelNames: nil,
		},
		{
			Type:       "variable",
			LabelNames: []string{"name"},
		},
		{
			Type:       "validation",
			LabelNames: nil,
		},
		{
			Type:       "features",
			LabelNames: nil,
		},
		{
			Type:       "locals",
			LabelNames: nil,
		},
		{
			Type:       "output",
			LabelNames: []string{"name"},
		},
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
		{
			Type:       "resource",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "data",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "module",
			LabelNames: []string{"name"},
		},
	},
}

// Variable block schema
var variableBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "description",
		},
		{
			Name: "default",
		},
		{
			Name: "type",
		},
		{
			Name: "sensitive",
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "validation",
		},
	},
}

var variableValidationBlockSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "validation",
		},
	},
	Attributes: []hcl.AttributeSchema{
		{
			Name: "condition",
		},
		{
			Name: "error_message",
		},
	},
}

// terraform block schema
var terraformBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "required_version",
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "required_providers",
		},
	},
}

// module block schema
var moduleConfigSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "source",
		},
		{
			Name: "version",
		},
		{
			Name: "providers",
		},
	},
}

// provider block schema
var providerConfigSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "version",
		},
		{
			Name: "alias",
		},
		{
			Name: "client_id",
		},
		{
			Name: "client_secret",
		},
		{
			Name: "subscription_id",
		},
		{
			Name: "tenant_id",
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "features",
		},
	},
}

// resource block schema
var resourceSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "provider",
		},
	},
}

// output block schema
var outputBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "description",
		},
		{
			Name: "value",
		},
		{
			Name: "depends_on",
		},
		{
			Name: "sensitive",
		},
	},
}
