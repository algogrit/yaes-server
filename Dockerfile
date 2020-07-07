FROM golang:1.14.2 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN make linux

FROM alpine:latest
RUN adduser -D non-root
USER non-root
WORKDIR /app
COPY --from=builder /app/yaes-server /app
EXPOSE 8000
EXPOSE 8080
ENTRYPOINT ./yaes-server
