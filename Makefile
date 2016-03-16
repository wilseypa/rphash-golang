install:
	go install github.com/wilseypa/rphash-golang/parse
	go install github.com/wilseypa/rphash-golang/types
	go install github.com/wilseypa/rphash-golang/utils
	go install github.com/wilseypa/rphash-golang/decoder
	go install github.com/wilseypa/rphash-golang/hash
	go install github.com/wilseypa/rphash-golang/itemset
	go install github.com/wilseypa/rphash-golang/clusterer
	go install github.com/wilseypa/rphash-golang/lsh
	go install github.com/wilseypa/rphash-golang/projector
	go install github.com/wilseypa/rphash-golang/reader
	go install github.com/wilseypa/rphash-golang/defaults
	go install github.com/wilseypa/rphash-golang/simple
	go install github.com/wilseypa/rphash-golang/stream

get:
	go get github.com/wilseypa/rphash-golang

all: install
