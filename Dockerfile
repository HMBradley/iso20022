FROM golang:1.15-alpine as builder
WORKDIR /go/src/github.com/moov-io/iso20022
RUN apk add -U make
COPY . .
RUN make build
