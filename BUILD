load("@rules_go//go:def.bzl", "go_library", "go_test")
load("@gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/therealfakemoot/wikilink-obsidian-resolver
gazelle(name = "gazelle")

go_library(
    name = "wikilink-obsidian-resolver",
    srcs = [
        "resolver.go",
        "version.go",
    ],
    importpath = "github.com/therealfakemoot/wikilink-obsidian-resolver",
    visibility = ["//visibility:public"],
    x_defs = {
        "github.com/therealfakemoot/wikilink-obsidian-resolver/resolver.Version": "{STABLE_STAMP_VERSION}",
        "github.com/therealfakemoot/wikilink-obsidian-resolver/resolver.Build": "{STABLE_STAMP_BUILD}",
    },
    deps = ["@dev_abhg_go_goldmark_wikilink//:wikilink"],
)

go_test(
    name = "wikilink-obsidian-resolver_test",
    srcs = ["resolver_test.go"],
    data = glob(["testdata/**"]),
    deps = [
        ":wikilink-obsidian-resolver",
        "@com_github_stretchr_testify//assert",
        "@dev_abhg_go_goldmark_wikilink//:wikilink",
    ],
)
