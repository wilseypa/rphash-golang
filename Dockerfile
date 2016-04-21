# grog dockerIp 172.17.0.2
# docker build -t rphash .
# docker run -i -t rphash /bin/bash

FROM golang:latest
RUN go get github.com/chrislusf/glow \
           github.com/wilseypa/rphash-golang/demo

EXPOSE 8080
