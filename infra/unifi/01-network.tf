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

resource "unifi_network" "iot" {
  name    = "iot"
  purpose = "corporate"

  subnet                       = "172.16.0.0/24"
  vlan_id                      = 2
  dhcp_start                   = "172.16.0.6"
  dhcp_stop                    = "172.16.0.254"
  dhcp_enabled                 = true
  internet_access_enabled      = false
  intra_network_access_enabled = false
}

resource "unifi_wlan" "iot-wifi" {
  name            = "iot-wifi"
  passphrase      = data.sops_file.unifi-secret.data["unifi.iotpw"]
  security        = "wpapsk"
  wpa3_support    = true
  wpa3_transition = true
  pmf_mode        = "optional"
  network_id      = unifi_network.iot.id
  ap_group_ids    = [data.unifi_ap_group.default.id]
  user_group_id   = data.unifi_user_group.default.id
  hide_ssid       = true
}
