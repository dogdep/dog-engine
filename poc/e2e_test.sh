#!/bin/bash
# Usage: ./e2e_test.sh

set -o errexit
set -o nounset
set -o pipefail

APP_NAME="fluffy_app"
APP_DEPLOYMENT="deployments/fluffy.tar.gz"
APP_IP_PORT="127.0.0.1:8080"

# initialization
make init
make build-base

# app deployment
pushd fluffy
bash dog.sh deploy

# app build and run on remote
popd
./build_and_run.sh ${APP_DEPLOYMENT} ${APP_NAME}

# Test whether its reacheable
sleep 10s
ret=`curl -I -s -o /dev/null -w "%{http_code}\n" "http://${APP_IP_PORT}" || true`

if [ "$ret" -ne "200" ]; then
  echo "[FAIL] Expected to see 200 OK on ${APP_IP_PORT}, however got {$ret}"
  docker stop $(docker ps -a -q --filter="name=${APP_NAME}")
  docker rm ${APP_NAME}
  exit 1
else
  echo "[OK] The app is responive at ${APP_IP_PORT}. Cleaning up containers".
  docker stop $(docker ps -a -q --filter="name=${APP_NAME}")
  docker rm ${APP_NAME}
  exit 0
fi
