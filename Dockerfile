FROM golang:1.19.1-alpine3.15 as builder
RUN apk add --no-cache alpine-sdk=1.0-r1
WORKDIR /build
COPY [ ".", "." ]
RUN make build
