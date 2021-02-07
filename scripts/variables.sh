#!/bin/bash

export CLIENT_VERSION="$(cat VERSION)"
export GIT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
export GO_VERSION="$(go version)"
