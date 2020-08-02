#!/usr/bin/env bash

kubectl delete namespace monitoring
kubectl delete namespace ingress

kubectl delete pvc --all
kubectl delete pv --all
