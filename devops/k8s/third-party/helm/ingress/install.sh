#!/usr/bin/env bash

helm install ingress --namespace ingress -f devops/k8s/third-party/helm/ingress/values.yaml nginx-stable/nginx-ingress
