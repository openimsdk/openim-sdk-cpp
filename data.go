package main

type BoolValue struct {
	Value bool `json:"value"`
}

type IntValue struct {
	Value int `json:"value"`
}

type StringValue struct {
	Value string `json:"value"`
}

type Empty struct {
}

type Error struct {
	OperationId string `json:"operationId"`
	ErrCode     int32  `json:"errCode"`
	ErrMsg      string `json:"errMsg"`
}
type Success struct {
	OperationId string `json:"operationId"`
	Data        string `json:"data"`
	DataType    int    `json:"dataType"`
}
type ErrorOrSuccess struct {
	OperationId string `json:"operationId"`
	ErrCode     int32  `json:"errCode"`
	Data        string `json:"data"`
	DataType    int    `json:"dataType"`
	ErrMsg      string `json:"errMsg"`
}

type Progress struct {
	OperationId string `json:"operationId"`
	Progress    int    `json:"progress"`
}

type MsgIDAndList struct {
	Id   string `json:"msgId"`
	List string `json:"list"`
}
