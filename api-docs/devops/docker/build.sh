#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

docker build -f $DIR/../../Dockerfile -t gauravagarwalr/yaes-api-docs:latest .
docker push gauravagarwalr/yaes-api-docs:latest
