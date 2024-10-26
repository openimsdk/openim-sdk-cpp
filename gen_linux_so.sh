#!/bin/bash
go build -buildmode=c-shared -trimpath -ldflags="-s -w" -o ./linux/libopenimsdk.so ./
