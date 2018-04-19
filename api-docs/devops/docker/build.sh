#!/usr/bin/env bash

ACTUAL_WD=$PWD
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd $DIR/../..

docker build -t gauravagarwalr/yaes-api-docs:latest .
docker push gauravagarwalr/yaes-api-docs:latest

cd $ACTUAL_WD
