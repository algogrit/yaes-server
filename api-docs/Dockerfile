FROM golang:alpine
RUN apk add --no-cache git mercurial \
    && go get -v -u github.com/go-swagger/go-swagger/cmd/swagger \
    && apk del git mercurial
COPY swagger.yml .
EXPOSE 3000
CMD swagger serve --no-open -p $PORT swagger.yml
