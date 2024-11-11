#!/bin/bash

pushd cmd
    go build
popd

pushd add-plugin
    go build -buildmode=plugin -o plugin.so plugin.go
popd

pushd sub-plugin
    go build -buildmode=plugin -o plugin.so plugin.go
popd

pushd mul-plugin
    go build -buildmode=plugin -o plugin.so plugin.go
popd

pushd div-plugin
    go build -buildmode=plugin -o plugin.so plugin.go
popd