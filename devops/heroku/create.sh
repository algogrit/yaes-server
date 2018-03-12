#!/usr/bin/env bash

heroku create yaes-server
heroku addons:create heroku-postgresql:hobby-dev --app yaes-server
heroku git:remote --app yaes-server
git push heroku master

heroku create yaes-server-dev
heroku addons:create heroku-postgresql:hobby-dev --app yaes-server-dev
heroku git:remote --app yaes-server-dev
git push heroku master

heroku create yaes-server-docker
heroku addons:create heroku-postgresql:hobby-dev --app yaes-server-docker
heroku git:remote --app yaes-server-docker
heroku container:push web --app yaes-server-docker
