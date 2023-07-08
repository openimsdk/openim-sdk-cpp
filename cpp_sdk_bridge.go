package main

import (
	"fmt"
	"open_im_sdk/open_im_sdk"
	"open_im_sdk/open_im_sdk_callback"
)

// #cgo LDFLAGS: -L. -lbridge
// #include <listeners_bridge.h>
import "C"

type OnConnListenerFromCpp struct {
}

func (l *OnConnListenerFromCpp) OnConnecting() {
	C.OnConnListenerOnConnecting()
}

func (l *OnConnListenerFromCpp) OnConnectSuccess() {
	C.OnConnListenerOnConnectSuccess()
}

func (l *OnConnListenerFromCpp) OnConnectFailed(errCode int32, errMsg string) {
	C.OnConnListenerOnConnectFailed(C.int(errCode), C.CString(errMsg))
}

func (l *OnConnListenerFromCpp) OnKickedOffline() {
	C.OnConnListenerOnKickedOffline()
}

func (l *OnConnListenerFromCpp) OnUserTokenExpired() {
	C.OnConnListenerOnUserTokenExpired()
}

func NewOnConnListener() open_im_sdk_callback.OnConnListener {
	return &OnConnListenerFromCpp{}
}

//export InitSDK
func InitSDK(operationID string, config string) {
	open_im_sdk.InitSDK(NewOnConnListener(), operationID, config)
}

func main() {
	fmt.Println("Should do nothing")
}
