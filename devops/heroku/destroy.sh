#!/usr/bin/env bash

heroku destroy yaes-server-docker --confirm yaes-server-docker
heroku destroy yaes-server-dev --confirm yaes-server-dev
heroku destroy yaes-server --confirm yaes-server
