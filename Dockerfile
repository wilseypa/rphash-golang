FROM golang:latest

RUN \
  go get github.com/wilseypa/rphash-golang && \
  go get github.com/chrislusf/glow
