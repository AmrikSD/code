load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load", "oci_push")

oci_image(
    name = "image",
    base = "@grafana",
)

oci_load(
    name = "load",
    image = ":image",
    repo_tags = ["grafana"]
)

oci_push(
    name = "push",
    image = ":image",
    repository = "ghcr.io/amriksd/grafana",
    remote_tags = ["latest"]
)

