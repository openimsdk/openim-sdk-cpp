#!/bin/sh
export LANG=en_US.UTF-8
# ios device
export CFLAGS="-arch arm64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphoneos --show-sdk-path) 
export CGO_LDFLAGS="-arch arm64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphoneos --show-sdk-path)  
export CGO_ENABLED=1 
export GOARCH=arm64 
export GOOS=ios
export CC="clang $CFLAGS $CGO_LDFLAGS" 
go build -tags ios -ldflags=-w -trimpath -v -o libopenimsdk.a -buildmode c-archive ./
# go build -tags ios -ldflags=-w -trimpath -v -o libopenimsdk_ios.a -buildmode c-archive ./
# ios simulator
# export CFLAGS="-arch x86_64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphonesimulator --show-sdk-path) 
# export CGO_LDFLAGS="-arch x86_64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphonesimulator --show-sdk-path) 
# CGO_ENABLED=1
# GOARCH=amd64 
# GOOS=darwin 
# CC="clang $CFLAGS $CGO_LDFLAGS" 
# go build -tags ios -ldflags=-w -trimpath -v -o libopenimsdk_iossimulator.a -buildmode c-archive ./
 
# lipo -create libopenimsdk_ios.a libopenimsdk_iossimulator.a -output libopenimsdk.a
