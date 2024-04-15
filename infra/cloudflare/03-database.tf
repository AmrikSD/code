resource "cloudflare_d1_database" "wedding" {
    account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
    name              = "wedding-database"
}
