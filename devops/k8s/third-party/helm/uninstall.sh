#!/usr/bin/env bash

set -e

./devops/k8s/third-party/helm/prometheus/uninstall.sh
./devops/k8s/third-party/helm/dashboard/uninstall.sh
./devops/k8s/third-party/helm/postgres/uninstall.sh
./devops/k8s/third-party/helm/ingress/uninstall.sh
