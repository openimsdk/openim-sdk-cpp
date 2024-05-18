#!/bin/bash

rm ./openimsdk.so ./openimsdk.h
go build -buildmode=c-shared -trimpath -ldflags="-s -w" -o openimsdk.so export.go constant.go protocol.go tools.go


# build cpp sdk
rm ./openimsdkcc.so
g++ -fPIC -shared -o openimsdkcc.so openimsdkcc.cc ./openimsdk.so