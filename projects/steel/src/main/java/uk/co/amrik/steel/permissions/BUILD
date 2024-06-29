load("@rules_java//java:defs.bzl", "java_library")

package(default_visibility = ["//visibility:public"])

filegroup(
    name = "permissions_root",
    srcs = glob(["*.java"]),
)

java_library(
    name = "permissions",
    srcs = [
        ":permissions_root",
    ],
    deps = [
        "//projects/steel/src/main/java/uk/co/amrik/steel/persistence",
        "@maven//:com_google_inject_guice",
        "@maven//:org_jooq_jooq"
    ]
)
