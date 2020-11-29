#!/bin/sh

echo Downloading dependencies
go get -t ./...

echo Building application
cd cmd/gochess && GOOS=linux GOARCH=amd64 go build -race -o ~/gochess || exit

echo Copying config
cp /go/src/github.com/vehsamrak/gochess/configs/config-dev.yml ~/config.yml

echo Running application
cd ~ && ./gochess
