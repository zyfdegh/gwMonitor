#!/bin/bash

echo "get packages..."
go get github.com/emicklei/go-restful
go get github.com/jmoiron/jsonq
go get github.com/gambol99/go-marathon
echo "get packages finished"

echo "build..."
go build -o monitor
echo "build finished"
