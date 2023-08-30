package main

/*
#include <stdio.h>

typedef void (*base_func)();
typedef void (*err_func)(int,void *);
typedef void (*success_func)(char *);


extern void c_base_caller(base_func func);
extern void c_err_caller(err_func func,int ,void*);
extern void c_success_caller(success_func func,char* data);


*/
import "C"
import (
	"open_im_sdk/open_im_sdk"
	"unsafe"
)

//export  init_sdk
func init_sdk(onConnecting C.base_func,
	onConnectSuccess C.base_func,
	onKickedOffline C.base_func,
	onUserTokenExpired C.base_func,
	onConnectFailed C.err_func,
	operationID *C.char, config *C.char) bool {
	callback := NewConnCallback(onConnecting, onConnectSuccess, onKickedOffline, onUserTokenExpired, onConnectFailed)
	return open_im_sdk.InitSDK(callback, C.GoString(operationID), C.GoString(config))
}

//export  login
func login(successFunc C.success_func, failedFunc C.err_func, operationID, uid, token *C.char) {
	baseCallback := NewBaseCallback(successFunc, failedFunc)
	open_im_sdk.Login(baseCallback, C.GoString(operationID), C.GoString(uid), C.GoString(token))
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
	C.c_base_caller(c.onConnecting)
}

func (c ConnCallback) OnConnectSuccess() {
	C.c_base_caller(c.onConnectSuccess)
}

func (c ConnCallback) OnConnectFailed(errCode int32, errMsg string) {
	C.c_err_caller(c.onConnectFailed, C.int(errCode), unsafe.Pointer(C.CString(errMsg)))

}

func (c ConnCallback) OnKickedOffline() {
	C.c_base_caller(c.onKickedOffline)
}

func (c ConnCallback) OnUserTokenExpired() {
	C.c_base_caller(c.onUserTokenExpired)
}

type BaseCallback struct {
	successFunc C.success_func
	failedFunc  C.err_func
}

func NewBaseCallback(successFunc C.success_func, failedFunc C.err_func) *BaseCallback {
	return &BaseCallback{successFunc: successFunc, failedFunc: failedFunc}
}

func (b BaseCallback) OnError(errCode int32, errMsg string) {
	C.c_err_caller(b.failedFunc, C.int(errCode), unsafe.Pointer(C.CString(errMsg)))
}

func (b BaseCallback) OnSuccess(data string) {
	C.c_success_caller(b.successFunc, C.CString(data))
}
