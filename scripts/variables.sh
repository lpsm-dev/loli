#!/bin/bash

CLIENT_VERSION="$(cat VERSION)"
GIT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
GO_VERSION="$(go version)"
export CLIENT_VERSION
export GIT_BRANCH
export GO_VERSION
