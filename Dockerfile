FROM golang:latest as builder
RUN mkdir -p /go/src/github.com/gauravagarwalr/yaes-server
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/gauravagarwalr/yaes-server
ADD . /go/src/github.com/gauravagarwalr/yaes-server
RUN rm -rf /go/src/github.com/gauravagarwalr/yaes-server/vendor
RUN dep ensure
RUN make linux

FROM alpine:latest
RUN adduser -D non-root
USER non-root
WORKDIR /app
COPY --from=builder /go/src/github.com/gauravagarwalr/yaes-server/yaes-server /app
EXPOSE 12345
ENTRYPOINT ./yaes-server
