load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_spf13_cobra",
        importpath = "github.com/spf13/cobra",
        sum = "h1:X+jTBEBqF0bHN+9cSMgmfuvv2VHJ9ezmFNf9Y/XstYU=",
        version = "v1.5.0",
    )
