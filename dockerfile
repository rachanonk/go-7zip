FROM golang:1.17rc2-alpine3.14
COPY . /go/src
RUN apk update \
    && apk search p7zip-full \
    && apk search p7zip \
    && apk add --update --no-cache p7zip
WORKDIR /go/src/cmd
CMD [ "go", "build" ]
CMD [ "go", "run", "main.go" ]
