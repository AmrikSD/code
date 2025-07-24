filegroup(
    name = "resources",
    srcs = [
        "application.conf",
        ":migrations",
    ],
    visibility = ["//projects/Rolodex/src/main/java/uk/co/amrik/rolodex:__pkg__"],
)

filegroup(
    name = "migrations",
    srcs = glob(
        ["db/migration/*"],
    ),
    visibility = ["//projects/Rolodex/src/main/java/uk/co/amrik/rolodex/database:__pkg__"],
)
