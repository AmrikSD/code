load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load", "oci_push")

oci_image(
    name = "image",
    base = "@wireguard",
)

oci_load(
    name = "load",
    image = ":image",
    repo_tags = ["wireguard"],
)

oci_push(
    name = "push",
    image = ":image",
    remote_tags = ["latest"],
    repository = "ghcr.io/amriksd/wireguard",
)
