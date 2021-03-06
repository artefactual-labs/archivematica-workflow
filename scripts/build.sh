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
  'linux/amd64'
  'windows/amd64'
)

if [ ! -z ${1+x} ]; then
  if [ "${1}" == '--current' ]; then
    eval $(go tool dist env | grep GOOS)
    eval $(go tool dist env | grep GOARCH)
    BUILDS=("${GOOS}/${GOARCH}")
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
    tmpdir="$(mktemp -d)"
    echo "Building ${name}: version=${version} goos=${goos} goarch=${goarch}"
    env CGO_ENABLED=0 GOOS=${goos} GOARCH=${goarch} go build -ldflags="-X github.com/artefactual-labs/archivematica-workflow/version.VERSION=${version}" -o ${tmpdir}/${name} ${pkg}
    env BZIP2=-1 tar -C ${tmpdir} -cjf ${bin_dir}/${name}-${goos}-${goarch}.tar.bz2 ${name}
    rm -rf ${tmpdir}
  done
}

echo "Finding git tag..."
version=`git describe --tags`

build_binary "archivematica-workflow" ${version}
build_binary "archivematica-workflow-cli" ${version}
