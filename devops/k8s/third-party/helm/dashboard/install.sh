#!/usr/bin/env bash

kubectl create namespace kubernetes-dashboard

helm install dashboard --namespace kubernetes-dashboard -f devops/k8s/third-party/helm/dashboard/values.yaml kubernetes-dashboard/kubernetes-dashboard || echo "Already installed k8s dashboard..."

kubectl apply -f ./devops/k8s/third-party/dashboard/dashboard-adminuser.yaml
