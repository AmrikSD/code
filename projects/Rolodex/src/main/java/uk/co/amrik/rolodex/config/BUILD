load("@rules_java//java:defs.bzl", "java_binary", "java_library")

package(default_visibility = ["//visibility:public"])

java_library(
    name = "config",
    srcs = glob(
        ["*.java"],
    ),
    deps = [
        "@maven//:com_google_inject_guice",
        "@maven//:com_typesafe_config",
        "@maven//:org_slf4j_slf4j_api",
    ],
)
