#!/usr/bin/env bash

ACTUAL_WD=$PWD
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd $DIR/../..

heroku container:push web --app yaes-api-docs

cd $ACTUAL_WD
