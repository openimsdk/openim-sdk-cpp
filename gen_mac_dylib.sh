export GOOS=darwin
export CGO_ENABLED=1
export CC=clang
export GOARCH=arm64
go build -buildmode=c-shared -o ./mac/libopenimsdk_arm64.dylib ./

export GOARCH=amd64
go build -buildmode=c-shared -o ./mac/libopenimsdk_amd64.dylib ./
lipo -create -output libopenimsdk.dylib libopenimsdk_arm64.dylib libopenimsdk_amd64.dylib