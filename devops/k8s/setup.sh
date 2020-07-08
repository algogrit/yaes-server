#!/usr/bin/env bash

set -e

helm repo add stable https://kubernetes-charts.storage.googleapis.com/
helm repo add nginx-stable https://helm.nginx.com/stable
helm repo add bitnami https://charts.bitnami.com/bitnami

helm repo update

kubectl create namespace ingress
kubectl create namespace monitoring
