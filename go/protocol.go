package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
typedef void (*CB_S)(char *);
typedef void (*CB_I_S)(int,char *);
typedef void (*CB_S_I_S_S)(char *,int,char *,char *);
typedef void (*CB_S_I_S_S_I)(char *,int,char *,char *,int);

CB_S DebugPrint;

void Call_CB_S(CB_S func,char* data)
{
    if(func == NULL){
        return;
    }
    func(data);
}


void Call_CB_I_S(CB_I_S func,int event,char* data)
{
    if(func == NULL){
        return;
    }
    if (strcmp(data, "\"\"") == 0) {
       strcpy(data, "");
    }
    func(event,data);
    if (data != NULL && data[0] != '\0')
    {
        free(data);
    }
}
void Call_CB_S_I_S_S(CB_S_I_S_S func,char* operationID, int errCode,char* errMsg,char* data)
{
    if(func == NULL){
        return;
    }
   if (strcmp(data, "\"\"") == 0) {
       strcpy(data, "");
    }
   if (strlen(errMsg) != 0 && errMsg[strlen(errMsg) - 1] != '\0') {
       strncat(errMsg, "\0", 1);
   }
    func(operationID,errCode,errMsg,data);
    if (errMsg != NULL && errMsg[0] != '\0')
    {
        free(errMsg);
    }
    if (data != NULL && data[0] != '\0')
    {
        free(data);
    }
    if (operationID != NULL)
    {
        free(operationID);
    }
}
void Call_CB_S_I_S_S_I(CB_S_I_S_S_I func,char* operationID,int errCode,char* errMsg,char* data,int progress)
{
    if(func == NULL){
        return;
    }
    if(strcmp(data, "\"\"") == 0) {
       strcpy(data, "");
    }
    func(operationID,errCode,errMsg,data,progress);
    if (errMsg != NULL && errMsg[0] != '\0')
    {
        free(errMsg);
    }
    if (data != NULL && data[0] != '\0')
    {
        free(data);
    }
      if (operationID != NULL)
    {
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
