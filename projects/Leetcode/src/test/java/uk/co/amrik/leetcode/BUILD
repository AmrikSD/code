load("@rules_java//java:defs.bzl", "java_test")

package(default_visibility = ["//visibility:public"])

java_test(
    name = "LeetcodeTest",
    srcs = ["LeetcodeTest.java"],
    deps = [
        "@maven//:junit_junit"
    ]
)

filegroup(
    name = "srcs",
    srcs = glob(["**/*.java"]) + ["BUILD"],
    visibility = ["//projects/Leetcode:__pkg__"],
)
