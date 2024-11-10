@echo off
set NDK_PATH=D:\OpenHarmonySDK\12\native

set SO_NAME=libopenimsdk
set OUT_PATH=openharmony\

set CGO_ENABLED=1

set NM=%NDK_PATH%\llvm\bin\llvm-nm
set AR=%NDK_PATH%\llvm\bin\llvm-ar
set LD=%NDK_PATH%\llvm\bin\ld.lld
set BASE_FLAGS=--sysroot=%NDK_PATH%\sysroot -fdata-sections -ffunction-sections -funwind-tables -fstack-protector-strong -no-canonical-prefixes -fno-addrsig -Wformat -Werror=format-security  -D__MUSL__ -fPIC -MD -MT -MF

REM armeabi-v7a
set GOOS=android
set GOARCH=arm
set CXXFLAGS=--target=arm-linux-ohos -march=armv7-a -mfloat-abi-softfp -generic-armv7-a -mthumb %BASE_FLAGS%
set CFLAGS=--target=arm-linux-ohos %BASE_FLAGS%
set CC=%NDK_PATH%\llvm\bin\clang %CFLAGS% %BASE_FLAGS%
set CXX=%NDK_PATH%\llvm\bin\clang++ %CFLAGS% %BASE_FLAGS%
@REM  not support 
@REM go build -buildmode=c-shared  -trimpath -ldflags="-s -w" -o %OUT_PATH%armeabi-v7a\%SO_NAME%.so ./

REM arm64-v8a
set GOOS=android
set GOARCH=arm64
set CXXFLAGS=--target=aarch64-linux-ohos %BASE_FLAGS%
set CFLAGS=--target=aarch64-linux-ohos %BASE_FLAGS%
set CC=%NDK_PATH%\llvm\bin\clang %CFLAGS% %BASE_FLAGS%
set CXX=%NDK_PATH%\llvm\bin\clang++ %CFLAGS% %BASE_FLAGS%

go build -buildmode=c-shared  -trimpath -ldflags="-s -w" -o %OUT_PATH%arm64-v8a\%SO_NAME%.so ./

REM x86_64
set GOOS=android
set GOARCH=amd64
set CXXFLAGS=--target=x86_64-linux-ohos %BASE_FLAGS%
set CFLAGS=--target=x86_64-linux-ohos %BASE_FLAGS%
set CC=%NDK_PATH%\llvm\bin\clang %CFLAGS% %BASE_FLAGS%
set CXX=%NDK_PATH%\llvm\bin\clang++ %CFLAGS% %BASE_FLAGS%

@REM go build -x -v -buildmode=c-shared  -trimpath -ldflags="-s -w" -o %OUT_PATH%x86_64\%SO_NAME%.so ./
go build  -buildmode=c-shared  -trimpath -ldflags="-s -w" -o %OUT_PATH%x86_64\%SO_NAME%.so ./


