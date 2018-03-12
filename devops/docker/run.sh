#!/usr/bin/env bash

docker run -it --rm -p 3000:3000 -e GO_APP_ENV=production \
  -e PORT=3000 \
  -e DATABASE_URL="postgres://$USER@localhost:5432/$DB_NAME" \
  gauravagarwalr/yaes-server
