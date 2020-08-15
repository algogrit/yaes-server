.PHONY: setup-db recreate-db build run

create-all-db:
	createdb yaes-dev
	createdb yaes-test
	createdb yaes

setup-db:
	createdb ${DB_NAME}
	${MAKE} migrate

recreate-db:
	dropdb ${DB_NAME}
	createdb ${DB_NAME}
	${MAKE} migrate

setup: setup-db
	dep ensure

build:
	go build ./cmd/server

migrate:
	go run ./cmd/migration

linux:
	GOOS=linux CGO_ENABLED=0 go build ./cmd/server

prod-build: GO_APP_ENV = production
prod-build: build

run: PORT = 3000
run:
	go run ./cmd/server/main.go

prod-run: GO_APP_ENV = production
prod-run: run

dev-setup:
	go get github.com/codegangsta/gin

dev-run:
	gin --build ./cmd/server

setup-docs:
	go get -v -u github.com/go-swagger/go-swagger/cmd/swagger

docs:
	swagger serve api-docs/swagger.yml

test: DB_NAME = yaes-test
test: recreate-db
	go generate ./...
	GO_APP_ENV="test" go test -v ./...

bench: DB_NAME = yaes-test
bench: recreate-db
	go generate ./...
	GO_APP_ENV="test" go test -bench=. -v ./...

ci-test:
	go generate ./...
	GO_APP_ENV="test" go test -v ./...

k8s-deploy:
	./devops/k8s/setup.sh
	./devops/k8s/up.sh

k8s-remove:
	./devops/k8s/down.sh
	./devops/k8s/teardown.sh

k8s-reapply-svc:
	kubectl delete -f devops/k8s/service.yaml
	kubectl apply -f devops/k8s/service.yaml

k8s-deploy-dashboard:
	./devops/k8s/third-party/dashboard/up.sh

k8s-remove-dashboard:
	./devops/k8s/third-party/dashboard/down.sh

k8s-proxy:
	kubectl proxy

k8s-open-dashboard:
	open http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy

k8s-get-dashboard-token:
	./devops/k8s/third-party/dashboard/scripts/get-token.sh

# username: admin
# password: prom-operator
kctl-pf-grafana:
	kubectl port-forward svc/monitoring-grafana 20000:80 -n monitoring

kctl-pf-prometheus:
	kubectl port-forward svc/monitoring-prometheus-oper-prometheus 9090 -n monitoring

# export POSTGRES_PASSWORD=$(kubectl get secret --namespace default yaes-db-postgresql -o jsonpath="{.data.postgresql-password}" | base64 --decode)
# PGPASSWORD="$POSTGRES_PASSWORD" psql postgres://yaesuser@localhost:5433/yaes?sslmode=disable
kctl-pf-postgres:
	kubectl port-forward --namespace default svc/yaes-db-postgresql 5433:5432

kctl-logs-yaes:
	kubectl logs -f -l app=yaes
