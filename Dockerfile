FROM golang:1.14.2 as builder
ARG PROJECT_NAME=server
WORKDIR /app
RUN echo Building image for... $PROJECT_NAME
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o /run/app ./cmd/${PROJECT_NAME}

FROM alpine:latest
# FROM scratch
# COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
RUN adduser -D non-root
USER non-root
WORKDIR /app
COPY --from=builder /run/app /app
ENV GO_APP_ENV production
ENV PORT 8080
EXPOSE 8080
ENV DIAGNOSTICS_PORT 8000
EXPOSE 8000
ENTRYPOINT ./app
