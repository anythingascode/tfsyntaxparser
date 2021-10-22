terraform {
  required_providers {
    hashicups = {
      version = "0.2"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {
  alias = "test"
  username = "education"
  password = "test123"
}

provider "azurerm" {
  alias = "testazure"
  client_id = "education"
  client_secret = "test123"
  subscription_id = "testsub"
  tenant_id = "testtenant"
  features {}
}


module "psl1" {
  source = "./coffee1"
  version = "1.0.1"
  coffee_name = "Packer Spiced Latte"
}

module "psl2" {
  source = "./coffee2"
  version = "1.0.2"
  coffee_name = "Packer Spiced Latte"
}

module "psl3" {
  source = "./coffee3"
  version = "1.0.3"
  coffee_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.coffee
}

data "hashicups_order" "order" {
  id = 1
}

output "order" {
  value = data.hashicups_order.order
}

variable "location" {
  type = string
  description = "(Required) Specifies the supported Azure location for Diagnostic Setting module."
  default = "eastus"
}

variable "test" {
  type = string
  validation {
    condition     = var.test == ""
    error_message = "Value must not be \"nope\"."
  }
}

variable "validation" {
  validation {
    condition     = var.validation == 5
    error_message = "Must be five."
  }
  default = 5
}
