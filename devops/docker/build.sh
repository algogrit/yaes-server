#!/usr/bin/env bash

docker build --build-arg PROJECT_NAME=migration -t gauravagarwalr/yaes-migration:latest .
docker push gauravagarwalr/yaes-migration:latest

docker build -t gauravagarwalr/yaes-server:latest .
docker push gauravagarwalr/yaes-server:latest
