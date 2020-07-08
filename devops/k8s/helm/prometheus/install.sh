#!/usr/bin/env bash

helm install monitoring --namespace=monitoring  -f devops/k8s/helm/prometheus/values.yaml stable/prometheus-operator
