# grog dockerIp 172.17.0.2
# docker build -t rphash .
# docker run -i -t rphash /bin/bash

FROM golang:latest

RUN mkdir -p src/github.com/wilseypa/rphash-golang
RUN go get github.com/chrislusf/glow \
  && go get github.com/wilseypa/rphash-golang

EXPOSE 8080
