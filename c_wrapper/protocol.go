package main

/*
#include <stdio.h>
typedef void (*CB_I_S)(int,char *);
typedef void (*CB_I_S_S)(int,char *,char *);
typedef void (*CB_I_S_S_I)(int,char *,char *,int);

__attribute__((weak))
void Call_CB_I_S(CB_I_S func,int event,char* data)
{
    func(event,data);
}
__attribute__((weak))
void Call_CB_I_S_S(CB_I_S_S func,int errCode,char* errMsg,char* data)
{
    func(errCode,errMsg,data);
}
__attribute__((weak))
void Call_CB_I_S_S_I(CB_I_S_S_I func,int errCode,char* errMsg,char* data,int progress)
{
    func(errCode,errMsg,data,progress);
}
enum CONN_EVENT{
   CONNECTING,
   CONNECT_SUCCESS,
   CONNECT_FAILED,
   KICKED_OFFLINE,
   USER_TOKEN_EXPIRED

};
*/
import "C"
