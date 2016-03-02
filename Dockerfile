#
# RPHash Dockerfile
#
# https://github.com/wenkesj/rphash
#
FROM ubuntu
RUN \
  mkdir -p /goroot && \
  curl https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1 \
  mkdir -p /gopath

ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH
WORKDIR /gopath
CMD ["bash"]
