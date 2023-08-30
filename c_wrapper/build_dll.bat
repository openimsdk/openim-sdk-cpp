go build -buildmode=c-shared  -trimpath -ldflags="-s -w"  -o openIM.dll  export.go

