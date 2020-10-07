@echo off

set GoDevWork=%cd%\

set GOOS=linux
set GOARCH=amd64

echo "Build For rpcx-pb-server ..."

cd %GoDevWork%rpcx-pb-server\
go build -ldflags "-s -w" -o rpcx-pb-server

echo "--------- Build For rpcx-pb-server Success!"

echo "Build For rpcx-pb-client ..."

cd %GoDevWork%rpcx-pb-client\
go build -ldflags "-s -w" -o rpcx-pb-client

echo "--------- Build For rpcx-pb-client Success!"

cd ..

pause