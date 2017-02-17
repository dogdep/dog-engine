#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

deploy() {
  tar czf ../deployments/fluffy.tar.gz .
}

case "$1" in
  deploy)
    deploy
    ;;
  *)
    echo "Usage: $0 {deploy}"
esac
