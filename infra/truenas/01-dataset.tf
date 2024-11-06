resource "truenas_dataset" "test" {
  pool = var.pool_name
  name = "test"
  quota_bytes = 12 * 1024 * 1024 * 1024
  comments = "[Managed By Terraform] - a test dataset"
}
