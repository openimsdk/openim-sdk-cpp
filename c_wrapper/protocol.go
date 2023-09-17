package main

/*
#include <stdio.h>
#include <stdlib.h>
typedef void (*CB_S)(char *);
typedef void (*CB_I_S)(int,char *);
typedef void (*CB_S_I_S_S)(char *,int,char *,char *);
typedef void (*CB_S_I_S_S_I)(char *,int,char *,char *,int);

CB_S DebugPrint;

void Call_CB_S(CB_S func,char* data)
{
    if(func == NULL){
        printf("callback func is null\n");
        return;
    }
    func(data);
}


void Call_CB_I_S(CB_I_S func,int event,char* data)
{
    if(func == NULL){
        printf("callback func is null\n");
        return;
    }
    func(event,data);
    if (data != NULL && data[0] != '\0')
    {
        printf("this is not null data event\n");
        free(data);
    }
}
void Call_CB_S_I_S_S(CB_S_I_S_S func,char* operationID, int errCode,char* errMsg,char* data)
{
    if(func == NULL){
        printf("callback func is null\n");
        return;
    }
    func(operationID,errCode,errMsg,data);
    if (errMsg != NULL && errMsg[0] != '\0')
    {
        printf("this is not null errmsg\n");
        free(errMsg);
    }
    if (data != NULL && data[0] != '\0')
    {
        printf("this is not null data base,opid:%s,data:%s\n",operationID,data);
        free(data);
    }
    if (operationID != NULL)
    {
        printf("this is not null operationID:%s\n",operationID);
        free(operationID);
    }
}
void Call_CB_S_I_S_S_I(CB_S_I_S_S_I func,char* operationID,int errCode,char* errMsg,char* data,int progress)
{
    if(func == NULL){
        printf("callback func is null\n");
        return;
    }
    func(operationID,errCode,errMsg,data,progress);
    if (errMsg != NULL && errMsg[0] != '\0')
    {
        printf("this is not null errmsg\n");
        free(errMsg);
    }
    if (data != NULL && data[0] != '\0')
    {
        printf("this is not null data\n");
        free(data);
    }
      if (operationID != NULL)
    {
        printf("this is not null operationID\n");
        free(operationID);
    }
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
