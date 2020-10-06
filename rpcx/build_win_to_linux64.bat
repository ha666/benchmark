@echo off

set GoDevWork=%cd%\

set GOOS=linux
set GOARCH=amd64

echo "Build For rpcx-server ..."

cd %GoDevWork%rpcx-server\
go build -ldflags "-s -w" -o rpcx-server

echo "--------- Build For rpcx-server Success!"

echo "Build For rpcx-client ..."

cd %GoDevWork%rpcx-client\
go build -ldflags "-s -w" -o rpcx-client

echo "--------- Build For rpcx-client Success!"

cd ..

pause