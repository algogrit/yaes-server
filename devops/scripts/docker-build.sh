#!/usr/bin/env bash

docker build -f devops/Dockerfile -t gauravagarwalr/yaes-server:latest .
docker push gauravagarwalr/yaes-server:latest
