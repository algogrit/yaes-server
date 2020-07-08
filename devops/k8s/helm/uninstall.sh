#!/usr/bin/env bash

set -e

./devops/k8s/helm/postgres/uninstall.sh
./devops/k8s/helm/ingress/uninstall.sh
./devops/k8s/helm/prometheus/uninstall.sh
