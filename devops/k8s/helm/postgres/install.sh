#!/usr/bin/env bash

helm install yaes-db -f devops/k8s/helm/postgres/values.yaml bitnami/postgresql
