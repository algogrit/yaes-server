# YAES (Yet Another Expense Splitter)

[![CircleCI](https://circleci.com/gh/algogrit/yaes-server/tree/master.svg?style=svg)](https://circleci.com/gh/algogrit/yaes-server/tree/master)

Expense Splitting Web App ( similar to Splitwise):

Primary purpose of the web app is to keep track of your expenses, payables and receivables to individuals.

**Docs available at**: https://yaes-api-docs.herokuapp.com/docs

## Planned Features

- Add users to your app. You will eventually split expense with them.
- Add expense and it can be sharable among selected users.
- Visibility of payables and receivables per user.
- Visibility of total payables and receivables.
- Visibility of all the individual expenses involving you at once place.
- Settle up with any user.
- Delete any expense.

## Clean Architecture by [Robert Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

![Clean Architecture](assets/CleanArchitecture.jpg)

- **Dependency rule**: Source code dependencies can only point inwards

## Dev

### Setup

```bash
go get -u algogrit.com/yaes-server/cmd/yaes-server
cd $GOPATH/src/algogrit.com/yaes-server
make setup
```

### Run Server

```bash
make run
```

### Dev Run

```bash
make dev-setup
make dev-run
```

### Docs

```bash
make setup-docs

make docs
```

### Tests

```bash
DB_NAME="yaes-test" make setup-db # Only first time
make test
```
