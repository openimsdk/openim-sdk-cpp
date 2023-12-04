export GOOS=darwin
export GOARCH=arm64
export CGO_ENABLED=1
export CC=clang

go build -buildmode=c-shared -o libopenimsdk.dylib export.go constant.go protocol.go tools.go


