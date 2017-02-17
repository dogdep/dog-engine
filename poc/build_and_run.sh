#!/bin/bash
# Usage: ./build_and_run.sh deployments/fluffy.tar.gz fluffy

set -o errexit
set -o nounset
set -o pipefail

# Clean workspace before starting
mkdir -p workspace
rm -rf workspace/*

# extract source and build
if [ ! -n "$1" ]; then
  echo "Error. File '${1}' does not exist or not specified"
  exit 1;
fi;

tar zxf $1 -C workspace

docker build -t $2:latest workspace
docker run -p 127.0.0.1:8080:8080 --name $2 -d -i $2
