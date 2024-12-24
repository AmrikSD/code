load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load", "oci_push")

oci_image(
    name = "image",
    base = "@home-assistant",
)

oci_load(
    name = "load",
    image = ":image",
    repo_tags = ["home-assistant"],
)

oci_push(
    name = "push",
    image = ":image",
    remote_tags = ["latest"],
    repository = "ghcr.io/amriksd/home-assistant",
)
