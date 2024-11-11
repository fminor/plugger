#!/bin/bash

pushd cmd
    go build
popd

pushd add-plugin
    go build -buildmode=plugin -o plugin.so plugin.go
popd