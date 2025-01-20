load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_spf13_cobra",
        importpath = "github.com/spf13/cobra",
        sum = "h1:X+jTBEBqF0bHN+9cSMgmfuvv2VHJ9ezmFNf9Y/XstYU=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_otiai10_gosseract_v2",
        importpath = "github.com/otiai10/gosseract/v2",
        sum = "h1:G8AyBpXEeSlcq8TI85LH/pM5SXk8Djy2GEXisgyblRw=",
        version = "v2.4.1",
    )
