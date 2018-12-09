#!/bin/sh
set -xe

TAG=$1

CGO_ENABLED=0 GOOS=linux go build -o whoami .

docker build -t kanatohodets/whoami:$TAG .
docker push kanatohodets/whoami:$TAG
