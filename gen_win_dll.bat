set CGO_ENABLED=1
go build -buildmode=c-shared  -trimpath -ldflags="-s -w"  -o openimsdk.dll  ./
