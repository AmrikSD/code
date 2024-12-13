variable "ip_address" {
  description = "The IP address to use for the Cloudflare DNS record"
  type        = string
}

resource "cloudflare_record" "amrik_xyz" {
  zone_id = data.sops_file.cloudflare-secret.data["cloudflare.amrik.xyz.zone_id"]
  name    = "@"
  type    = "A"
  content   = var.ip_address
}
