package main

/*
#include<stdio.h>
typedef void (*MessageHandler)(int id ,char* data);
MessageHandler messageHandler;
void CallMessageHandler(MessageHandler msgHandler,int id,char* data){
    if(msgHandler == NULL){
        printf("IMSDK: not set message handler");
    }else{
        msgHandler(id,data);
    }
}

*/
import "C"
