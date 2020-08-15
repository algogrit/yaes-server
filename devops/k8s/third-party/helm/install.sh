#!/usr/bin/env bash

set -e

./devops/k8s/third-party/helm/prometheus/install.sh
./devops/k8s/third-party/helm/ingress/install.sh
./devops/k8s/third-party/helm/postgres/install.sh
