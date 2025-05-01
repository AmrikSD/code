variable "github_org" {
  type = string
}
variable "github_repository" {
  type = string
}

resource "github_repository" "this" {
  name               = var.github_repository
  allow_merge_commit = false
  allow_rebase_merge = false
  has_issues         = true
  has_downloads      = true
}

resource "tls_private_key" "flux" {
  algorithm   = "ECDSA"
  ecdsa_curve = "P256"
}

resource "github_repository_deploy_key" "this" {
  title      = "Flux [Managed By Terraform]"
  repository = github_repository.this.name
  key        = tls_private_key.flux.public_key_openssh
  read_only  = "false"
}

resource "flux_bootstrap_git" "this" {
  depends_on = [github_repository_deploy_key.this]
  count = 0

  embedded_manifests = true
  path               = "infra/kubernetes/cluster-02/"
}
