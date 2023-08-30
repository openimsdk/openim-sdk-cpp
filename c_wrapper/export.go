package main

/*
#include <stdio.h>
typedef void (*CB)();
typedef void (*CB_I_S)(int,char *);
typedef void (*CB_S)(char *);
__attribute__((weak))
void Call_CB(CB func)
{
    func();
}
__attribute__((weak))
void Call_CB_I_S(CB_I_S func,int errCode,char* errMsg)
{
    func(errCode,errMsg);
}
__attribute__((weak))
void Call_CB_S(CB_S func,char* data)
{
    func(data);
}
*/
import "C"

import (
    "open_im_sdk/open_im_sdk"
)

type ConnCallback struct {
    onConnecting       C.CB
    onConnectSuccess   C.CB
    onConnectFailed    C.CB_I_S
    onKickedOffline    C.CB
    onUserTokenExpired C.CB
}

func NewConnCallback(onConnecting C.CB, onConnectSuccess C.CB,
    onKickedOffline C.CB, onUserTokenExpired C.CB, onConnectFailed C.CB_I_S) *ConnCallback {
    return &ConnCallback{onConnecting: onConnecting, onConnectSuccess: onConnectSuccess,
        onKickedOffline: onKickedOffline, onUserTokenExpired: onUserTokenExpired, onConnectFailed: onConnectFailed}
}

func (c ConnCallback) OnConnecting() {
    C.Call_CB(c.onConnecting)
}

func (c ConnCallback) OnConnectSuccess() {
    C.Call_CB(c.onConnectSuccess)
}

func (c ConnCallback) OnConnectFailed(errCode int32, errMsg string) {
    C.Call_CB_I_S(c.onConnectFailed, C.int(errCode), C.CString(errMsg))

}

func (c ConnCallback) OnKickedOffline() {
    C.Call_CB(c.onKickedOffline)
}

func (c ConnCallback) OnUserTokenExpired() {
    C.Call_CB(c.onUserTokenExpired)
}

type BaseCallback struct {
    successFunc C.CB_S
    failedFunc  C.CB_I_S
}

func NewBaseCallback(successFunc C.CB_S, failedFunc C.CB_I_S) *BaseCallback {
    return &BaseCallback{successFunc: successFunc, failedFunc: failedFunc}
}

func (b BaseCallback) OnError(errCode int32, errMsg string) {
    C.Call_CB_I_S(b.failedFunc, C.int(errCode), C.CString(errMsg))
}

func (b BaseCallback) OnSuccess(data string) {
    C.Call_CB_S(b.successFunc, C.CString(data))
}

//export  init_sdk
func init_sdk(onConnecting C.CB,
    onConnectSuccess C.CB,
    onKickedOffline C.CB,
    onUserTokenExpired C.CB,
    onConnectFailed C.CB_I_S,
    operationID *C.char, config *C.char) bool {
    callback := NewConnCallback(onConnecting, onConnectSuccess, onKickedOffline, onUserTokenExpired, onConnectFailed)
    return open_im_sdk.InitSDK(callback, C.GoString(operationID), C.GoString(config))
}

//export  login
func login(successFunc C.CB_S, failedFunc C.CB_I_S, operationID, uid, token *C.char) {
    baseCallback := NewBaseCallback(successFunc, failedFunc)
    open_im_sdk.Login(baseCallback, C.GoString(operationID), C.GoString(uid), C.GoString(token))
}

func main() {

}
