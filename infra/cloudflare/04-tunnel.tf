resource "cloudflare_tunnel" "amrik_co_uk" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  name       = "amrik-co-uk"
  secret     = data.sops_file.cloudflare-secret.data["cloudflare.amrik.co.uk.tunnel.secret"]
}
