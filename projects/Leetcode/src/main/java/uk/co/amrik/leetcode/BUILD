load("@rules_java//java:defs.bzl", "java_binary", "java_library")

package(default_visibility = ["//visibility:public"])

java_library(
    name = "lc-library",
    srcs = glob(
        ["*.java"],
    ),
)

java_binary(
    name = "lc-binary",
    main_class = "uk.co.amrik.leetcode.Leetcode",
    runtime_deps = [":lc-library"]
)


filegroup(
    name = "srcs",
    srcs = ["BUILD"] + glob(["**/*.java"]),
)
