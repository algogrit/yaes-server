# Kubernetes installation

## Setup & getting the cluster up & running

### Setup helm, repos & namespaces

```bash
./devops/k8s/setup.sh
```

### Installing the charts and getting the services up & running

```bash
./devops/k8s/up.sh
```

This internally calls `./devops/k8s/helm/install.sh`.

## Clean the cluster

`down.sh` internally calls `./devops/k8s/helm/uninstall.sh`.

```bash
./devops/k8s/down.sh
./devops/k8s/teardown.sh
```
