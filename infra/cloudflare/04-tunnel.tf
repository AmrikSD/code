resource "cloudflare_tunnel" "amrik_co_uk" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  name       = "amrik-co-uk"
  secret     = data.sops_file.cloudflare-secret.data["cloudflare.amrik.co.uk.tunnel.secret"]
}

resource "cloudflare_tunnel_config" "example_config" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  tunnel_id  = cloudflare_tunnel.amrik_co_uk.id

  config {
    warp_routing {
      enabled = true
    }
    origin_request {
      connect_timeout          = "1m0s"
      tls_timeout              = "1m0s"
      tcp_keep_alive           = "1m0s"
      no_happy_eyeballs        = false
      keep_alive_connections   = 1024
      keep_alive_timeout       = "1m0s"
      http_host_header         = "baz"
      origin_server_name       = "foobar"
      ca_pool                  = "/path/to/unsigned/ca/pool"
      no_tls_verify            = false
      disable_chunked_encoding = false
      bastion_mode             = false
      proxy_address            = "10.0.0.1"
      proxy_port               = "8123"
      proxy_type               = "socks"
      ip_rules {
        prefix = "/web"
        ports  = [80, 443]
        allow  = false
      }
    }
    ingress_rule {
      hostname = "foo"
      path     = "/bar"
      service  = "http://10.0.0.2:8080"
      origin_request {
        connect_timeout = "2m0s"
        access {
          required  = true
          team_name = "terraform"
          aud_tag   = ["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"]
        }
      }
    }
    ingress_rule {
      service = "https://10.0.0.3:8081"
    }
  }
}
