#!/usr/bin/env bash

# This script is used to compile the protocol buffers.

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
__root="$(cd "$(dirname "${__dir}")" && pwd)"

proto_src="${__root}/proto"
proto_pkg="${__root}/pkg/service/pb"
proto_gopathsrc="$(cd "$(dirname "${__root}/../../../../")" && pwd)"

protoc --version >/dev/null 2>&1 || {
	echo >&2 "Protocol Buffers is not installed.";
	echo >&2 "See https://developers.google.com/protocol-buffers/ for instructions.";
	echo >&2 "Aborting.";
	exit 1;
}

rm ${proto_pkg}/*.pb.go || true

# If we ever use gogo/protobuf
# proto_gogosrc="$(cd "$(dirname "${__root}/../../../github.com/gogo/protobuf/protobuf")" && pwd)"
# cd ${proto_src} && protoc --proto_path=${proto_gopathsrc}:${proto_gogosrc}:./ --gogo_out=plugins=grpc:${proto_pkg} *.proto

# Vanilla protobuf
echo "Compiling..."
cd ${proto_src} && protoc --proto_path=${proto_gopathsrc}:./ --go_out=plugins=grpc:${proto_pkg} *.proto

ls ${proto_src}/*.proto
