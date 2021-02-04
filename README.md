# gRPC Server

## Quick Start

```bash
# install golang
$ tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz
$ export PATH=$PATH:/usr/local/go/bin

# set GOPATH
$ mkdir $HOME/go
$ export GOPATH=$HOME/go
$ export PATH=$GOPATH/bin:$PATH

# generate *.go from *.proto
$ ./progo-gen.sh

# install project dependencies
$ go get
$ go mod vendor

# build binary
$ make build

# run server
$ ./app

# run client
$ RUN_MODE=client ./app

# HTTP call
$ curl http://localhost:9080/healthcheck
{}
```
