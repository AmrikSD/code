load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_cross_binary" )
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load", "oci_push", "oci_image_index")
load("@rules_pkg//pkg:tar.bzl", "pkg_tar")

go_library(
    name = "app_lib",
    srcs = glob(["cmd/uttt/*.go"]),
    visibility = ["//visibility:private"],
)

go_binary(
    name = "app-linux-amd64",
    embed = [":app_lib"],
    visibility = ["//visibility:public"],
    goos = "linux",
    goarch = "arm64",
    out = "uttt-linux-amd64"
)

go_binary(
    name = "app-linux-arm64",
    embed = [":app_lib"],
    visibility = ["//visibility:public"],
    goos = "linux",
    goarch = "arm64",
    out = "uttt-linux-arm64"
)

pkg_tar(
    name = "tar-linux-arm64",
    srcs = [":app-linux-arm64"],
    symlinks = {
        "uttt": "uttt-linux-arm64"
    }
)

pkg_tar(
    name = "tar-linux-amd64",
    srcs = [":app-linux-amd64"],
    symlinks = {
        "uttt": "uttt-linux-amd64"
    }
)

oci_image(
    name = "image-linux-arm64",
    tars = [
        ":tar-linux-arm64"
    ],
    os = "linux",
    architecture = "arm64",
    entrypoint = ["./uttt"],
)

oci_image(
    name = "image-linux-amd64",
    tars = [
        ":tar-linux-amd64"
    ],
    os = "linux",
    architecture = "amd64",
    entrypoint = ["./uttt"],
)

oci_image_index(
    name = "index",
    images = [
       ":image-linux-arm64",
       ":image-linux-amd64"
    ]
)

oci_push(
    name = "push",
    image = ":index",
    repository = "ghcr.io/amriksd/uttt",
    remote_tags = [ "latest" ]
)
