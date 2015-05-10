#!/bin/bash
set -ex

main() {
    ./main > ./output/issues.json
    echo 'init(' $(cat ./output/issues.json) ')' > ./output/issues.js
    exit 0
}

main
