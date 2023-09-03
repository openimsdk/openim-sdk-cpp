package main

import "encoding/json"

func parseBool() bool {
	return false
}

func StructToJsonString(param interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}
