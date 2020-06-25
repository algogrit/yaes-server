#!/usr/bin/env bash

set -e

./devops/k8s/helm/install.sh

kubectl apply -f devops/k8s/secrets.yaml
kubectl apply -f devops/k8s/service.yaml
