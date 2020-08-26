#!/usr/bin/env bash

set -e

helm install yaes-db -f devops/k8s/third-party/helm/postgres/values.yaml bitnami/postgresql || echo "Already installed postgres db..."
# helm upgrade yaes-db -f devops/k8s/third-party/helm/postgres/values.yaml bitnami/postgresql

kubectl wait \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/instance=yaes-db \
  --timeout=120s
