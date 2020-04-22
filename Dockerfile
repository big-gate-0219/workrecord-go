FROM golang:1.14-alpine3.11

SHELL ["/bin/ash", "-c"]
ENV GO111MODULE=on

WORKDIR /go/src/app
COPY ./src ./

RUN apk add --no-cache alpine-sdk
RUN go get github.com/pilu/fresh

EXPOSE 8080

ENTRYPOINT [ "fresh" ]
