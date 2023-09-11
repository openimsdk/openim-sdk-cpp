package main

/*
#include <stdio.h>
#include <stdlib.h>
typedef void (*CB_I_S)(int,char *);
typedef void (*CB_S_I_S_S)(char *,int,char *,char *);
typedef void (*CB_S_I_S_S_I)(char *,int,char *,char *,int);

void Call_CB_I_S(CB_I_S func,int event,char* data)
{
    func(event,data);
    free(data);
}
void Call_CB_S_I_S_S(CB_S_I_S_S func,char* operationID, int errCode,char* errMsg,char* data)
{
    func(operationID,errCode,errMsg,data);
    free(errMsg);
    free(data);
}
void Call_CB_S_I_S_S_I(CB_S_I_S_S_I func,char* operationID,int errCode,char* errMsg,char* data,int progress)
{
    func(operationID,errCode,errMsg,data,progress);
    free(errMsg);
    free(data);
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
