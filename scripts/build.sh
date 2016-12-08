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
  'linux/386'
  'linux/amd64'
  'linux/arm'
  'linux/arm64'
  'windows/386'
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

for BUILD in "${BUILDS[@]}"; do
   IFS='/' read -ra BUILD <<< "$BUILD"
   goos=${BUILD[0]}
   goarch=${BUILD[1]}

   exe="archivematica-workflow-${goos}-${goarch}"
   pkg="./cmd/archivematica-workflow"
   echo "Building ${exe} binary..."
   env CGO_ENABLED=0 GOOS=${goos} GOARCH=${goarch} go build -o ${bin_dir}/${exe} ${pkg}
done

for BUILD in "${BUILDS[@]}"; do
   IFS='/' read -ra BUILD <<< "$BUILD"
   goos=${BUILD[0]}
   goarch=${BUILD[1]}

   exe="archivematica-workflow-cli-${goos}-${goarch}"
   pkg="./cmd/archivematica-workflow-cli"
   echo "Building ${exe} binary..."
   env CGO_ENABLED=0 GOOS=${goos} GOARCH=${goarch} go build -o ${bin_dir}/${exe} ${pkg}
done
