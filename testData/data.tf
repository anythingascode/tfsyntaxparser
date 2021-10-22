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
}
