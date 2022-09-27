#!/bin/sh
set -e

rm -rf completions
mkdir completions

for sh in bash zsh fish; do
  go run ./cmd/loli/main.go completion "$sh" >"completions/loli.$sh"
done
