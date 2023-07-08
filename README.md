# openim-sdk-cpp

This branch is only for experimenting the communications between Go and C++. 

listeners_bridge.* solely are enough, but listeners.* can add a layer of encapsulation for users.

To export the Go API calls, firstly we need to compile the C++ interfaces imeplemented by user so Go can use them:

```
make
```

Then Go can use the provided interfaces to export API calls into C shared libraries:

```
 go build -o sdk.so -buildmode=c-shared cpp_sdk_bridge.go
```
