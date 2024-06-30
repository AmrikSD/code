load("@rules_java//java:defs.bzl", "java_binary", "java_library", "java_test")

load("@bazel_tools//tools/jdk:default_java_toolchain.bzl", "default_java_toolchain")
default_java_toolchain(
    name = "my_java_toolchain",
    source_version = "21",
    target_version = "21",
)

package(default_visibility = ["//visibility:public"])
