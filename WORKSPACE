load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "musl_toolchains",
    sha256 = "79c3dab39eb52a9c0927f6a222e2a10b733fd875ccdad59feadd5b30fd4518fb",
    url = "https://github.com/bazel-contrib/musl-toolchain/releases/download/v0.1.9/musl_toolchain-v0.1.9.tar.gz",
)

load("@musl_toolchains//:repositories.bzl", "load_musl_toolchains")

load_musl_toolchains()

load("@musl_toolchains//:toolchains.bzl", "register_musl_toolchains")

register_musl_toolchains()
