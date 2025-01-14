variable "shellys" {
    type = list(
    object({
        mac = string
        name = string
    })
    )
}

resource "unifi_user" "shellys" {
  for_each = zipmap(range(length(var.shellys)), var.shellys)

  name = each.value.name
  mac  = each.value.mac
  note = "Managed By Terraform"

  fixed_ip = format("172.16.0.%d", each.key + 10) # Start incrementing from 172.16.0.10

  network_id = unifi_network.iot.id
}

resource "unifi_firewall_group" "shellys" {
    name = "shellys"
    type = "address-group"
    members =  [ for k, shelly in unifi_user.shellys : shelly.fixed_ip]
}
