#!/usr/bin/env bash

GIT_REVISION=$(git rev-parse --short --verify HEAD)
TIME=$(date -u +%Y%m%d.%H%M%S)
VERSION=1.0.${GIT_REVISION}.${TIME}


go build -ldflags "-X main.versionString=${VERSION}" ./cmd/crypta/
