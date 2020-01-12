#!/usr/bin/env bash

set -e

# Warm up server
curl 'http://localhost:3000/'

# POST /users - Creating primary user algogrit
curl --location --request POST 'http://localhost:3000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
  "Password": "password",
  "FirstName": "G",
  "LastName": "A",
  "MobileNumber": "+916666666666",
  "Username": "algogrit"
}
'

# POST /users - Creating other user testuser
curl --location --request POST 'http://localhost:3000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
  "Password": "password",
  "FirstName": "Test",
  "LastName": "User",
  "MobileNumber": "+911234567890",
  "Username": "testuser"
}
'

# POST /login
curl --location --request POST 'http://localhost:3000/login' \
--header 'Content-Type: application/json' \
--data-raw '{
	"Password": "password",
	"Username": "algogrit"
}' | tee /tmp/yaes-login.json

# Grab the token
_token=`cat /tmp/yaes-login.json | jq -r .token`

# GET /users
curl --location --request GET 'http://localhost:3000/users' \
--header "Authorization: Bearer $_token" | jq .
