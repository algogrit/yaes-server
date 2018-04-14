#!/usr/bin/env bash

heroku create yaes-server
heroku addons:create heroku-postgresql:hobby-dev --app yaes-server
heroku git:remote --app yaes-server
heroku config:set SENTRY_DSN=$SENTRY_DSN --app yaes-server
heroku config:set SENTRY_RELEASE=production --app yaes-server
heroku config:set GO_APP_ENV=production --app yaes-server
git push heroku master

heroku create yaes-server-dev
heroku addons:create heroku-postgresql:hobby-dev --app yaes-server-dev
heroku git:remote --app yaes-server-dev
heroku config:set SENTRY_DSN=$SENTRY_DSN --app yaes-server-dev
heroku config:set SENTRY_RELEASE=staging --app yaes-server-dev
heroku config:set GO_APP_ENV=staging --app yaes-server-dev
git push heroku master
