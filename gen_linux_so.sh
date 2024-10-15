#!/bin/bash
rm ./openimsdk.so ./openimsdk.h
go build -buildmode=c-shared -trimpath -ldflags="-s -w" -o openimsdk.so ./
