#!/usr/bin/env bash

kubectl apply -f devops/k8s/third-party/helm/prometheus/monitoring-ns.yaml

helm install monitoring --namespace=monitoring  -f devops/k8s/third-party/helm/prometheus/values.yaml stable/prometheus-operator || echo "Already installed prometheus monitoring..."
