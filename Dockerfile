# grog dockerIp 172.17.0.2
# docker build -t rphash .
#
FROM golang:latest

RUN mkdir -p src/github.com/wilseypa
RUN git clone https://github.com/wilseypa/rphash-golang src/github.com/wilseypa
RUN go get github.com/chrislusf/glow
RUN curl -Lks https://bintray.com$(curl -Lk http://bintray.com/chrislusf/seaweedfs/seaweedfs/_latestVersion | grep linux_amd64.tar.gz | sed -n "/href/ s/.*href=['\"]\([^'\"]*\)['\"].*/\1/gp") | gunzip | tar -xf - -C /bin/weed
RUN chmod +x /bin/weed
