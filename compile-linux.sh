#!/bin/bash

go build ./main.go
go build ./launcher-linux.go

mv ./main ./towebm
