load("@rules_java//java:defs.bzl", "java_library")

package(default_visibility = ["//visibility:public"])

filegroup(
    name = "persistence_root",
    srcs = glob(["*.java"]),
)

java_library(
    name = "persistence",
    srcs = [
        ":persistence_root",
    ],
    resources = ["//projects/steel/src/main/resources:migration_files"],
    javacopts = ["-source", "21", "-target", "21"],
    deps = [
        "@maven//:com_google_inject_guice",
        "@maven//:com_google_inject_extensions_guice_servlet",
        "@maven//:org_flywaydb_flyway_core",
        "@maven//:org_flywaydb_flyway_database_postgresql",
        "@maven//:org_postgresql_postgresql",
        "@maven//:org_jooq_jooq",
    ]
)
