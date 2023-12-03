package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"strconv"
	"unsafe"
)

func parseBool(b int) bool {
	return !(b == 0)
}

func StructToJsonString(param interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func FreeCString(strList ...*C.char) {
	for _, str := range strList {
		C.free(unsafe.Pointer(str))
	}
}
func Int32ToString(intValue int32) string {
	return strconv.Itoa(int(intValue))
}
