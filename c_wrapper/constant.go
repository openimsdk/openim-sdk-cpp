package main

/*
#include <stdio.h>
*/
import "C"

var (
	NO_ERR      = C.int(0)
	NO_ERR_MSG  = C.CString("")
	NO_DATA     = C.CString("")
	NO_PROGRESS = C.int(0)
)

const (
	CONNECTING = iota
	CONNECT_SUCCESS
	CONNECT_FAILED
	KICKED_OFFLINE
	USER_TOKEN_EXPIRED
)
