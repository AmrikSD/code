variable "shellys" {
  type = list(
    object({
      mac  = string
      name = string
    })
  )
}

resource "unifi_user" "shellys" {
  for_each = { for idx, shelly in var.shellys : idx => shelly }

  name = each.value.name
  mac  = each.value.mac
  note = "Managed By Terraform"

  fixed_ip         = format("10.0.220.%d", each.key + 110) # Start incrementing from 10.0.220.110
  local_dns_record = format("shelly-%d.int.amrik.co.uk", each.key + 110)

  network_id = unifi_network.iot.id
}

resource "unifi_firewall_group" "shellys" {
  name    = "shellys"
  type    = "address-group"
  members = [for k, shelly in unifi_user.shellys : shelly.fixed_ip]
}
