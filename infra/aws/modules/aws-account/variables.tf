variable name {
    type = string
}

variable email {
    type = string
}

variable region {
    type = string
}

variable "tags" {
  description = "Tags to apply to the account and resources"
  type        = map(string)
  default     = {}
}

variable "access_key" {
    type = string
    sensitive = true
}

variable "secret_key" {
    type = string
    sensitive = true
}
