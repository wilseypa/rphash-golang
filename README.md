# Scalable Big Data Clustering by Random Projection Hashing #
+ Sam Wenke, Jacob Franklin, Sadiq Quasem
+ Advised by Dr. Phillip A. Wilsey

## Installing and Using Go ##
+ [Install Go](https://golang.org/doc/install?download=go1.5.1.windows-amd64.msi#uninstall)

```bash
# Export the go path.
export GOPATH=$HOME/rphash && cd $GOPATH
```

## Building ##
```bash
go get github.com/wenkesj/rphash
go build github.com/wenkesj/rphash
go test "github.com/wenkesj/rphash"
```

## Development Testing ##
```bash
go test
```

## API ##
