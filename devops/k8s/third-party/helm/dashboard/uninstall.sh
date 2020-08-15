#!/usr/bin/env bash

kubectl delete -f ./devops/k8s/third-party/dashboard/dashboard-adminuser.yaml

helm uninstall dashboard -n kubernetes-dashboard

kubectl delete namespace kubernetes-dashboard
