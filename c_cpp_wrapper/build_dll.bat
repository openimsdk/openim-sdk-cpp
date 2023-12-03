go build -buildmode=c-shared  -trimpath -ldflags="-s -w"  -o openimsdk.dll  export.go constant.go protocol.go tools.go

