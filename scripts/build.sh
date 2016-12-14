#!/usr/bin/env bash

# This script is used to build the Go sources.

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
__root="$(cd "$(dirname "${__dir}")" && pwd)"

bin_dir="${__root}/bin"
mkdir -p ${bin_dir}

cd ${__root}

# Full list: `go tool dist list`
BUILDS=(
  'darwin/amd64'
  'freebsd/amd64'
  'linux/amd64'
  'linux/arm'
  'linux/arm64'
  'windows/amd64'
)

if [ ! -z ${1+x} ]; then
  if [ "${1}" == '--current' ]; then
    eval $(go tool dist env | grep GOOS)
    eval $(go tool dist env | grep GOARCH)
    BUILDS=("${GOOS}/${GOARCH}")
  elif [ "${1}" == '--linux-amd64' ]; then
    BUILDS=("linux/amd64")
  fi
fi

function build_binary {
  name="${1}"
  version="${2}"
  pkg="./cmd/${name}"
  for BUILD in "${BUILDS[@]}"; do
    IFS='/' read -ra BUILD <<< "$BUILD"
    goos=${BUILD[0]}
    goarch=${BUILD[1]}
    bin="${name}-${version}-${goos}-${goarch}"
    echo "Building ${bin} binary..."
    env CGO_ENABLED=0 GOOS=${goos} GOARCH=${goarch} go build -ldflags="-X github.com/artefactual-labs/archivematica-workflow/pkg/version.VERSION=${version}" -o ${bin_dir}/${bin} ${pkg}
  done
}

echo "Finding git tag..."
version=`git describe --tags`

build_binary "archivematica-workflow" ${version}
build_binary "archivematica-workflow-cli" ${version}
