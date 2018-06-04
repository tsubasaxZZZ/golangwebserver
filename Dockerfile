FROM golang:1.8.3-alpine
MAINTAINER Tsubasa Nomura

ENV webserver_path /go/src/github.com/tsubasaxZZZ/golangwebserver/
ENV PATH $PATH:$webserver_path

WORKDIR $webserver_path
COPY main.go .

RUN go build .

ENTRYPOINT ./golangwebserver

EXPOSE 80
