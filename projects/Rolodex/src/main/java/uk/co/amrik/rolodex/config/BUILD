load("@rules_java//java:defs.bzl", "java_binary", "java_library")

java_library(
    name = "config",
    srcs = glob(
        ["*.java"],
    ),
    visibility = [
        "//projects/Rolodex/src/main/java/uk/co/amrik/rolodex:__pkg__",
        "//projects/Rolodex/src/main/java/uk/co/amrik/rolodex/database:__pkg__",
        "//projects/Rolodex/src/main/java/uk/co/amrik/rolodex/services:__pkg__",
    ],
    deps = [
        "@maven//:com_google_inject_guice",
        "@maven//:com_typesafe_config",
        "@maven//:org_slf4j_slf4j_api",
    ],
)
