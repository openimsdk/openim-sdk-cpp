# openim-sdk-cpp

Based on the secondary closure of [openim-sdk-core](https://github.com/openimsdk/openim-sdk-core), export the library of the corresponding platform

# golang 编译 openharmany so 代码修改

1. runtime/cgo/cgo.go
   #cgo android LDFLAGS: -llog
   改为  
    #cgo android LDFLAGS: -lhilog_ndk.z.so

2. runtime/cgo/gcc_android.c

#include <android/log.h>
改为
#include <hilog/log.h>

\_\_android_log_vprint(ANDROID_LOG_FATAL, "runtime/cgo", format, ap);
改为
OH_LOG_FATAL(LOG_APP, format, ap);

3. net/cgo_resold.go 参数错误修改

```golang
func cgoNameinfoPTR(b []byte, sa *C.struct_sockaddr, salen C.socklen_t) (int, error) {
	// gerrno, err := C.getnameinfo(sa, salen, (*C.char)(unsafe.Pointer(&b[0])), C.size_t(len(b)), nil, 0, C.NI_NAMEREQD)
	// openharmany
	gerrno, err := C.getnameinfo(sa, salen, (*C.char)(unsafe.Pointer(&b[0])), C.uint(len(b)), nil, 0, C.NI_NAMEREQD)
	return int(gerrno), err
}
```

# scripts

```bash
gen_win_dll.bat # gen windows .dll
gen_android_so.bat #  gen android .so
gen_ios_dylib.sh # gen ios .a
gen_linux_so.sh # gen linux .so
gen_mac_dylib.sh # gen mac .dylib
```
