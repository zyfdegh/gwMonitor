#!/bin/bash

echo "change direction"
cd ../../
echo $(pwd)

echo "export path"
export GOPATH=$(pwd)
echo "GOPATH:"$GOPATH

echo "get packages..."
go get github.com/emicklei/go-restful
go get github.com/jmoiron/jsonq
echo "get packages finished"

echo "build..."
CGO_ENABLED=0 go build -a -installsuffix cgo monitor
echo "build finished"
