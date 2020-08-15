#!/usr/bin/env bash

kubectl -n kubernetes-dashboard describe secret  $(kubectl -n kubernetes-dashboard get secret | grep admin-user | awk '{print $1}')  | ag 'token:' | cut -f 7 -d ' '
