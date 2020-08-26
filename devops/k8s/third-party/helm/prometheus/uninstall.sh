#!/usr/bin/env bash

helm uninstall monitoring -n monitoring

kubectl delete -f devops/k8s/third-party/helm/prometheus/monitoring-ns.yaml
