FROM golang:latest as builder
RUN go get -v -u github.com/go-swagger/go-swagger/cmd/swagger
COPY swagger.yml .
EXPOSE 3000
CMD swagger serve --no-open -p $PORT swagger.yml
