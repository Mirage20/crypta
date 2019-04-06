#!/usr/bin/env bash

GIT_REVISION=$(git rev-parse --short --verify HEAD)
TIME=$(date -u +%Y%m%d.%H%M%S)
VERSION=1.0.${GIT_REVISION}.${TIME}

build_artifacts () {
  local os=$1
  local arch=$2
  GOOS=$os GOARCH=$arch go build -ldflags "-X main.versionString=${VERSION}" ./cmd/crypta/
  file crypta
  tar -czvf crypta-darwin-x64.tar.gz crypta
  rm crypta
}

build_artifacts linux amd64
build_artifacts darwin amd64
