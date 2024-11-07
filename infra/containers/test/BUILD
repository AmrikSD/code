load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load", "oci_push")
load("@rules_pkg//pkg:tar.bzl", "pkg_tar")

filegroup(
    name = "static",
    srcs = ["index.html", "textfile.txt"],
)

pkg_tar(
    name = "static_tar",
    srcs = ["//infra/containers/test:static"],
    package_dir = "/usr/share/nginx/html"
)

oci_image(
    name = "image",
    base = "@nginx",  # Reference the pulled nginx image
    tars = [
        ":static_tar"
    ]
)

oci_load(
    name = "load",
    image = ":image",
    repo_tags = ["test:test"]
)

oci_push(
    name = "push",
    image = ":image",
    repository = "ghcr.io/amriksd/test",
    remote_tags = ["latest"]
)

