#!/usr/bin/env bash

set -e

kubectl delete -f devops/k8s/monitor.yaml
kubectl delete -f devops/k8s/service.yaml

./devops/k8s/helm/uninstall.sh
