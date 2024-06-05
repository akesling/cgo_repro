#!/usr/bin/env bash
set -e

_utils_dir="$(pwd)/$(dirname "$0")"
_repo_root="${_utils_dir}/.."
_dockerfile_dir="${_repo_root}/docker"

_build_staging_dir="${_repo_root}/build/docker/"
_dockerfile="${_repo_root}/Dockerfile"

_build_target="//repro:linux_tarball"

echo "Creating build staging directory ${_build_staging_dir}"
mkdir -p "${_build_staging_dir}"

echo "Building tarball"
# Build straight to print status / progress
bazel build -c opt "${_build_target}"
# Run again to capture file name
_release_tarball="$(bazel build -c opt "${_build_target}" 2>&1 | grep "/tar.tar" | tr -d ' ')"

echo "Copying dependencies to build tree for use by docker"
rsync -a --delete "${_repo_root}/${_release_tarball}" "${_build_staging_dir}/repro.tar"

echo "Building docker container"
docker build -f "${_dockerfile}" "${_build_staging_dir}" -t "repro:latest"
