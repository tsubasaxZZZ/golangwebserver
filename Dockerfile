FROM golang:1.8.3-alpine
MAINTAINER Tsubasa Nomura

COPY main.go .

RUN go build .

ENTRYPOINT ./golangwebserver

EXPOSE 80
