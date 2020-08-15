#!/usr/bin/env bash

helm uninstall yaes-db

kubectl delete pvc -l app.kubernetes.io/instance=yaes-db
