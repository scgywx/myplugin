#!/bin/bash
projectpath=`pwd`
export GOPATH="${projectpath}:$GOPATH"
export GOBIN="${projectpath}/bin"
go build -o engine engine
