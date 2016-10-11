#!/bin/bash

# copy template
cp pgw.json ../../bin/pgw.json

echo "change direction"
cd ../../
echo $(pwd)

echo "export path"
export GOPATH=$(pwd)
echo "GOPATH:"$GOPATH

echo "get packages..."
go get github.com/emicklei/go-restful
go get github.com/jmoiron/jsonq
go get github.com/gambol99/go-marathon
echo "get packages finished"

echo "build..."
CGO_ENABLED=0 go build -o bin/monitor -a -installsuffix cgo monitor
echo "build finished"
