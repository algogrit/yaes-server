.PHONY: setup-db recreate-db build run

create-all-db:
	createdb yaes-dev
	createdb yaes-test
	createdb yaes

setup-db:
	createdb ${DB_NAME}

recreate-db:
	dropdb ${DB_NAME}
	createdb ${DB_NAME}

setup: setup-db
	dep ensure

build:
	go build

linux:
	GOOS=linux CGO_ENABLED=0 go build

prod-build: GO_APP_ENV = production
prod-build: build

run: PORT = 3000
run:
	go run main.go

prod-run: GO_APP_ENV = production
prod-run: run

dev-setup:
	go get github.com/codegangsta/gin

dev-run:
	gin

setup-docs:
	go get -v -u github.com/go-swagger/go-swagger/cmd/swagger

docs:
	swagger serve swagger.yml

test: DB_NAME = yaes-test
test: recreate-db
	GO_APP_ENV="test" go test -v ./...
