load("@rules_java//java:defs.bzl", "java_binary", "java_library")

package(default_visibility = ["//visibility:public"])

java_library(
    name = "services",
    srcs = glob(
        ["*.java"],
    ),
    deps = [
        "@maven//:com_google_inject_guice",
        "@maven//:org_slf4j_slf4j_api",
    ],
)

java_library(
    name = "health",
    srcs = glob(
        [
            "health/*.java",
        ],
    ),
    deps = [
        "@maven//:jakarta_servlet_jakarta_servlet_api",
    ],
)

java_library(
    name = "watch",
    srcs = glob(
        [
            "BackgroundService.java",
            "watch/*.java",
        ],
    ),
    deps = [
        "//projects/Rolodex/src/main/java/uk/co/amrik/rolodex/config",
        "@maven//:com_google_inject_guice",
        "@maven//:com_typesafe_config",
        "@maven//:org_slf4j_slf4j_api",
    ],
)
