#!/usr/bin/env bash

docker run -it --rm -p 3000:3000 -p 8080:8080\
  -e GO_APP_ENV=$GO_APP_ENV \
  -e PORT=3000 \
  -e DIAGNOSTICS_PORT=8080 \
  -e DATABASE_URL="postgres://$USER@docker.for.mac.host.internal/$DB_NAME?sslmode=disable" \
  gauravagarwalr/yaes-server
