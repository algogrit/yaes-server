#!/usr/bin/env bash

heroku git:remote --app yaes-server
git push heroku master

heroku git:remote --app yaes-server-dev
git push heroku master

heroku container:push web --app yaes-server-docker

git remote remove heroku
