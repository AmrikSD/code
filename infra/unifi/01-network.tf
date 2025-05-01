data "unifi_ap_group" "default" {
}

data "unifi_user_group" "default" {
}

resource "unifi_network" "default" {
  name    = "Default"
  purpose = "corporate"

  subnet      = "10.0.0.0/14"
  vlan_id     = 0
  domain_name = "int.amrik.co.uk"

  dhcp_enabled = true
  dhcp_dns = [
    "1.1.1.1",
    "9.9.9.9",
    "8.8.8.8",
    "8.8.4.4",
  ]
  dhcp_start             = "10.0.0.56"
  dhcp_stop              = "10.3.255.254"
  dhcp_v6_start          = "::2"
  dhcp_v6_stop           = "::7d1"
  ipv6_pd_start          = "::2"
  ipv6_pd_stop           = "::7d1"
  ipv6_ra_enable         = true
  ipv6_ra_priority       = "high"
  ipv6_ra_valid_lifetime = 0
}
