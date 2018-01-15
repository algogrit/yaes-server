.PHONY: setup-db build

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
  glide install

build:
  go build

run: build
  ./Yet-Another-Expense-Splitter

dev-run:
  gin

setup-docs:
  go get -v -u github.com/go-swagger/go-swagger/cmd/swagger

docs:
  swagger serve swagger.yml
