#!/usr/bin/env bash

kubectl create namespace ingress

helm install ingress --namespace ingress -f devops/k8s/third-party/helm/ingress/values.yaml nginx-stable/nginx-ingress || echo "Already installed nginx ingress..."
