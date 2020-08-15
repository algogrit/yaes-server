#!/usr/bin/env bash

helm uninstall ingress -n ingress
kubectl delete namespace ingress
