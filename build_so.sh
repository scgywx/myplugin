#!/bin/bash
projectpath=`pwd`
export GOPATH="${projectpath}:$GOPATH"
export GOBIN="${projectpath}/bin"
pluginname="plugin$1"
pluginpath="plugin_version_$1"
go build -buildmode plugin -o $pluginname.so --ldflags="-pluginpath=$pluginpath" logic
