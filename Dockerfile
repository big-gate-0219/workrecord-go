FROM golang:1.14-alpine3.11 AS builder
SHELL ["/bin/ash", "-c"]
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY ./src ./
RUN apk add --no-cache alpine-sdk
RUN go build server.go

FROM alpine:latest
SHELL ["/bin/ash", "-c"]
RUN apk add --no-cache alpine-sdk
WORKDIR /workrecord-go/
COPY --from=builder /go/src/app/server .
COPY --from=builder /go/src/app/.env .
ENTRYPOINT ["./server"]
EXPOSE 8080
