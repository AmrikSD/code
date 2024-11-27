variable "pool" {
  type = string
}

variable "datasets" {
  type = list(object({
    name        = string
    size_gb     = optional(number)
    description = optional(string)
  }))
}

resource "truenas_dataset" "datasets" {
  for_each    = { for ds in var.datasets : ds.name => ds }
  pool        = var.pool
  name        = each.key
  quota_bytes = each.value.size_gb == null ? 0 : each.value.size_gb * 1024 * 1024 * 1024
  comments    = each.value.description != null ? format("[Managed By Terraform] - %s", each.value.description) : "[Managed By Terraform]"
}
