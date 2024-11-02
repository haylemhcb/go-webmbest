#!/bin/bash

#!/bin/bash


export CGO_ENABLED=1
export GOARCH=amd64
export GOOS=windows
export CXX=x86_64-w64-mingw32-g++
export CC=x86_64-w64-mingw32-gcc

go build -ldflags '-s -w' ./towebm-windows.go
go build -ldflags '-s -w'  ./launcher-windows.go

