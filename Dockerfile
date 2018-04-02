FROM golang:latest as builder
RUN mkdir -p /go/src/github.com/gauravagarwalr/Yet-Another-Expense-Splitter
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/gauravagarwalr/Yet-Another-Expense-Splitter
ADD . /go/src/github.com/gauravagarwalr/Yet-Another-Expense-Splitter
RUN rm -rf /go/src/github.com/gauravagarwalr/Yet-Another-Expense-Splitter/vendor
RUN dep ensure
RUN make linux

FROM alpine:latest
RUN adduser -D non-root
USER non-root
WORKDIR /app
COPY --from=builder /go/src/github.com/gauravagarwalr/Yet-Another-Expense-Splitter/Yet-Another-Expense-Splitter /app
EXPOSE 12345
ENTRYPOINT ./Yet-Another-Expense-Splitter
