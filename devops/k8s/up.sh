#!/usr/bin/env bash

set -e

./devops/k8s/third-party/helm/install.sh

kubectl apply -f devops/k8s/service.yaml
kubectl apply -f devops/k8s/monitor.yaml
