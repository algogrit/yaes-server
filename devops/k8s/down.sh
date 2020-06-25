#!/usr/bin/env bash

set -e

kubectl delete -f devops/k8s/service.yaml
kubectl delete -f devops/k8s/secrets.yaml

./devops/k8s/helm/uninstall.sh