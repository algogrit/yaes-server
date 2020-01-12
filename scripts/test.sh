#!/usr/bin/env bash

# make create-db # if first run!
make recreate-db dev-setup

echo 'Started server? `$ make dev-run`'
read

./scripts/test/api.sh
