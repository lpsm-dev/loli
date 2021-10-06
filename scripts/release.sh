#!/bin/bash

npm run release-debug
git commit -am "chore: bump version" && git push
npm run release && git pull && git fetch
make release
