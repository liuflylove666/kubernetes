package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "clientset_generated.go",
        "doc.go",
        "register.go",
    ],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/serializer:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/watch:go_default_library",
        "//vendor/k8s.io/client-go/discovery:go_default_library",
        "//vendor/k8s.io/client-go/discovery/fake:go_default_library",
        "//vendor/k8s.io/client-go/testing:go_default_library",
        "//vendor/k8s.io/metrics/pkg/apis/metrics/v1alpha1:go_default_library",
        "//vendor/k8s.io/metrics/pkg/apis/metrics/v1beta1:go_default_library",
        "//vendor/k8s.io/metrics/pkg/client/clientset_generated/clientset:go_default_library",
        "//vendor/k8s.io/metrics/pkg/client/clientset_generated/clientset/typed/metrics/v1alpha1:go_default_library",
        "//vendor/k8s.io/metrics/pkg/client/clientset_generated/clientset/typed/metrics/v1alpha1/fake:go_default_library",
        "//vendor/k8s.io/metrics/pkg/client/clientset_generated/clientset/typed/metrics/v1beta1:go_default_library",
        "//vendor/k8s.io/metrics/pkg/client/clientset_generated/clientset/typed/metrics/v1beta1/fake:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)
