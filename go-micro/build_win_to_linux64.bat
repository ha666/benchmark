@echo off

set GoDevWork=%cd%\

set GOOS=linux
set GOARCH=amd64

echo "Build For micro-server ..."

cd %GoDevWork%micro-server\
go build -ldflags "-s -w" -o micro-server

echo "--------- Build For micro-server Success!"

echo "Build For micro-client ..."

cd %GoDevWork%micro-client\
go build -ldflags "-s -w" -o micro-client

echo "--------- Build For micro-client Success!"

cd ..

pause