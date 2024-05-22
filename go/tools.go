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

func IntToString[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](intValue T) string {
	return strconv.FormatInt(int64(intValue), 10)
}
