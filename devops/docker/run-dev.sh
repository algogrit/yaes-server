#!/usr/bin/env bash

docker run -it --rm -p 3000:3000 \
  -e GO_APP_ENV=$GO_APP_ENV \
  -e PORT=3000 \
  -e DATABASE_URL="postgres://$USER@docker.for.mac.host.internal/$DB_NAME?sslmode=disable" \
  gauravagarwalr/yet-another-expense-splitter
