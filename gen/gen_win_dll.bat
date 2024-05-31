set CGO_ENABLED=1
go build -buildmode=c-shared  -trimpath -ldflags="-s -w"  -o openimsdk.dll  export.go protocol.go message.go



@REM g++ -shared -fPIC -o openimsdkcc.dll openimsdkcc.cc openimsdk.dll