FROM golang:latest as builder

MAINTAINER Gaurav Agarwal <gaurav@tarkalabs.com>

RUN mkdir -p /go/src/github.com/gauravagarwalr/Yet-Another-Expense-Splitter
ADD . /go/src/github.com/gauravagarwalr/Yet-Another-Expense-Splitter
WORKDIR /go/src/github.com/gauravagarwalr/Yet-Another-Expense-Splitter
RUN make build

FROM alpine:latest
COPY --from=builder /go/src/github.com/gauravagarwalr/Yet-Another-Expense-Splitter/Yet-Another-Expense-Splitter .
EXPOSE 12345
WORKDIR /
CMD ./Yet-Another-Expense-Splitter
