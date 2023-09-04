set NDK_PATH=D:\android_sdk\ndk-bundle
set SO_NAME=libopenimsdk
set OUT_PATH=android\

set CGO_ENABLED=1

REM 生成 armeabi-v7a
set GOOS=android
set GOARCH=arm
set CC=%NDK_PATH%\toolchains\llvm\prebuilt\windows-x86_64\bin\armv7a-linux-androideabi16-clang.cmd

go build -buildmode=c-shared  -trimpath -ldflags="-s -w" -o %OUT_PATH%armeabi-v7a\%SO_NAME%.so export.go constant.go protocol.go tools.go

REM 生成 arm64-v8a
set GOARCH=arm64
set CC=%NDK_PATH%\toolchains\llvm\prebuilt\windows-x86_64\bin\aarch64-linux-android21-clang.cmd

go build -buildmode=c-shared  -trimpath -ldflags="-s -w" -o %OUT_PATH%arm64-v8a\%SO_NAME%.so export.go constant.go protocol.go tools.go

REM 生成 x86
set GOARCH=386
set CC=%NDK_PATH%\toolchains\llvm\prebuilt\windows-x86_64\bin\i686-linux-android16-clang.cmd

go build -buildmode=c-shared  -trimpath -ldflags="-s -w" -o %OUT_PATH%x86\%SO_NAME%.so export.go constant.go protocol.go tools.go

REM 生成 x86_64
set GOARCH=amd64
set CC=%NDK_PATH%\toolchains\llvm\prebuilt\windows-x86_64\bin\x86_64-linux-android21-clang.cmd

go build -buildmode=c-shared  -trimpath -ldflags="-s -w" -o %OUT_PATH%x86_64\%SO_NAME%.so export.go constant.go protocol.go tools.go


