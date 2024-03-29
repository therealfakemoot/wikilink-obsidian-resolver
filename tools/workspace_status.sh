#! /usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

echo "STABLE_STAMP_VERSION $(git describe --tags --dirty=-dev)"
echo "STABLE_STAMP_COMMIT $(git rev-parse HEAD)"
echo "STABLE_STAMP_BRANCH $(git rev-parse --abbrev-ref HEAD)"
