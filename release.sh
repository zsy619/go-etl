#!/bin/bash

sh cover.sh

export IGNORE_PACKAGES=db2
export CGO_ENABLED=1
export GO111MODULE=on
go mod download
go mod vendor
go generate ./...
cd cmd/datax

# 【darwin/amd64】
echo "start build darwin/amd64 ..."
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o datax-darwin
cp -f datax-darwin ../../release/datax-darwin
chmod -R 777 ../../release/datax-darwin
rm -f datax-darwin

# 【linux/amd64】
# echo "start build linux/amd64 ..."
# CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o datax-linux
# cp -f datax-linux ../../release/datax-linux
# chmod -R 777 ../../release/datax-linux
# rm -f datax-linux

# 【windows/amd64】
echo "start build windows/amd64 ..."
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 go build -v -o datax.exe
cp -f datax.exe ../../release/datax.exe
chmod -R 777 ../../release/datax.exe
rm -f datax.exe

cd ../..
go run tools/datax/release/main.go