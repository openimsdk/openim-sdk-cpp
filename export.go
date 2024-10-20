package main

/*
#include <stdio.h>
#include <stdlib.h>
typedef void (*MessageHandler)(int id ,char* data);
extern MessageHandler messageHandler;
extern void CallMessageHandler(MessageHandler msgHandler,int id,char* data);
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

//export set_msg_handler_func
func set_msg_handler_func(handler C.MessageHandler) {
	C.messageHandler = handler
}

//export call_api
func call_api(_apiKey C.int, jsonArgsStr *C.char) *C.char {
	apiKey := int(_apiKey)
	args := C.GoString(jsonArgsStr)
	res, err := callAPI(apiKey, args)
	if err != nil {
		DispatorMsg(Msg_Error, "CallAPI:"+err.Error())
		return C.CString("")
	}
	return C.CString(res)
}

// Actively release functions with return values ​​to avoid memory leakage
//
//export free_data
func free_data(p *C.char) {
	C.free(unsafe.Pointer(p))
}

func DispatorMsg(msgId int, msg interface{}) {
	t := reflect.TypeOf(msg)
	kind := t.Kind()
	var data = ""
	if kind == reflect.Struct {
		msgJson, err := json.Marshal(msg)
		if err != nil {
			msgId = 0
			data = fmt.Sprintf("Marshal Json Error :%s", err.Error())
		} else {
			data = string(msgJson)
		}
	} else if kind == reflect.String {
		data = msg.(string)
	} else if kind == reflect.Int32 {
		data = strconv.Itoa(int(msg.(int32)))
	}
	var cdata = C.CString(data)
	C.CallMessageHandler(C.messageHandler, C.int(msgId), cdata)
	C.free(unsafe.Pointer(cdata))
}

func callAPI(apiKey int, jsonArgs string) (string, error) {
	apiFunc, ok := GetAPIFunc(apiKey)
	if ok {
		return apiFunc(jsonArgs)
	} else {
		return "", fmt.Errorf("not find api key:%d,please check sdk", apiKey)
	}
}

func main() {

}
