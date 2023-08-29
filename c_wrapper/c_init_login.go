package main

/*
#include <stdio.h>

typedef void (*base_func)();
typedef void (*err_func)(int,void *);

extern base_func _onConnecting;
extern base_func _onConnectSuccess;
extern base_func _onKickedOffline;
extern base_func _onUserTokenExpired;
extern err_func _onConnectFailed;

extern void c_onConnecting();
extern void c_onConnectSuccess();
extern void c_onKickedOffline();
extern void c_onUserTokenExpired();
extern void c_onConnectFailed(int ,void*);


*/
import "C"
import (
	"open_im_sdk/open_im_sdk"
	"unsafe"
)

//export  InitSDK
func InitSDK(onConnecting C.base_func,
	onConnectSuccess C.base_func,
	onKickedOffline C.base_func,
	onUserTokenExpired C.base_func,
	onConnectFailed C.err_func,
	operationID *C.char, config *C.char) bool {
	callback := NewConnCallback(onConnecting, onConnectSuccess, onKickedOffline, onUserTokenExpired, onConnectFailed)
	return open_im_sdk.InitSDK(callback, C.GoString(operationID), C.GoString(config))
}
func main() {

}

type ConnCallback struct {
	onConnecting       C.base_func
	onConnectSuccess   C.base_func
	onKickedOffline    C.base_func
	onUserTokenExpired C.base_func
	onConnectFailed    C.err_func
}

func NewConnCallback(onConnecting C.base_func, onConnectSuccess C.base_func,
	onKickedOffline C.base_func, onUserTokenExpired C.base_func, onConnectFailed C.err_func) *ConnCallback {
	return &ConnCallback{onConnecting: onConnecting, onConnectSuccess: onConnectSuccess,
		onKickedOffline: onKickedOffline, onUserTokenExpired: onUserTokenExpired, onConnectFailed: onConnectFailed}
}

func (c ConnCallback) OnConnecting() {
	C._onConnecting = c.onConnecting
	C.c_onConnecting()
}

func (c ConnCallback) OnConnectSuccess() {
	C._onConnectSuccess = c.onConnectSuccess
	C.c_onConnectSuccess()
}

func (c ConnCallback) OnConnectFailed(errCode int32, errMsg string) {
	C._onConnectFailed = c.onConnectFailed
	C.c_onConnectFailed(C.int(errCode), unsafe.Pointer(C.CString(errMsg)))
}

func (c ConnCallback) OnKickedOffline() {
	C._onKickedOffline = c.onKickedOffline
	C.c_onKickedOffline()
}

func (c ConnCallback) OnUserTokenExpired() {
	C._onUserTokenExpired = c.onUserTokenExpired
	C.c_onUserTokenExpired()
}
