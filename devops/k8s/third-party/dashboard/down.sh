#!/usr/bin/env bash

kubectl delete -f devops/k8s/third-party/dashboard/dashboard-adminuser.yaml
kubectl delete -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.3/aio/deploy/recommended.yaml
