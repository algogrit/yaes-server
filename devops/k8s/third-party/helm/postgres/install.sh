#!/usr/bin/env bash

helm install yaes-db -f devops/k8s/third-party/helm/postgres/values.yaml bitnami/postgresql
# helm upgrade yaes-db -f devops/k8s/third-party/helm/postgres/values.yaml bitnami/postgresql
