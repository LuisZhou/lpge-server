#!/bin/bash

export GOPATH=$PWD/
export GOBIN=$PWD/bin

if [ $# -eq 1 ] && [ "$1" != "doc" ] && [ "$1" != "compile" ] ; then
	go install $1
elif [ "$1" == "test" ] ; then
	go test ${@:3} $2
elif [ "$1" == "doc" ] ; then	
	godoc -http=:6060 ${@:2}
elif [ "$1" == "compile" ] ; then
	go install server
else
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/jinzhu/inflection
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/LuisZhou/lpge/
	go install server
fi
