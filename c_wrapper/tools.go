package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"unsafe"
)

func parseBool() bool {
	return false
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
