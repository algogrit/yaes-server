FROM golang:1.14.2 as builder
WORKDIR /go/src/algogrit.com/yaes-server
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN make linux

FROM alpine:latest
RUN adduser -D non-root
USER non-root
WORKDIR /app
COPY --from=builder /go/src/algogrit.com/yaes-server/yaes-server /app
EXPOSE 12345
ENTRYPOINT ./yaes-server
