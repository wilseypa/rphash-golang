# grog dockerIp 172.17.0.2
# docker build -t rphash .
#
FROM golang:latest

RUN mkdir -p src/github.com/wilseypa
RUN git clone https://github.com/wilseypa/rphash-golang src/github.com/wilseypa
RUN go get github.com/chrislusf/glow

EXPOSE 8080
