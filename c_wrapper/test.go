package main

/*
#include <stdio.h>
typedef void (*base_func)();
typedef void (*err_func)(int,void *);

base_func _onConnecting;
 base_func _onConnectSuccess;
base_func _onKickedOffline;
 base_func _onUserTokenExpired;
 err_func _onConnectFailed;

void c_onConnecting()
{
_onConnecting();
}
void c_onConnectSuccess()
{
_onConnectSuccess();
}
void c_onKickedOffline()
{
_onKickedOffline();
}
void c_onUserTokenExpired()
{
_onUserTokenExpired();
}
void c_onConnectFailed(int errCode,void* errMsg)
{
_onConnectFailed(errCode,errMsg);
}
*/
import "C"
