load("@rules_java//java:defs.bzl", "java_binary", "java_library")

java_library(
    name = "database",
    srcs = glob(
        ["*.java"],
    ),
    resources = [
        "//projects/Rolodex/src/main/resources:migrations",
    ],
    visibility = [
        "//projects/Rolodex/src/main/java/uk/co/amrik/rolodex:__pkg__",
    ],
    deps = [
        "//projects/Rolodex/src/main/java/uk/co/amrik/rolodex/config",
        "@maven//:com_google_inject_guice",
        "@maven//:com_typesafe_config",
        "@maven//:com_zaxxer_HikariCP",
        "@maven//:org_flywaydb_flyway_core",
        "@maven//:org_postgresql_postgresql",
        "@maven//:org_slf4j_slf4j_api",
    ],
)
