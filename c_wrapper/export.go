package main

/*
#include <stdio.h>
typedef void (*base_func)();
typedef void (*err_func)(int,void *);
typedef void (*success_func)(char *);

void c_base_caller(base_func func)
{
    func();
}
void c_err_caller(err_func func,int errCode,void* errMsg)
{
    func(errCode,errMsg);
}
void c_success_caller(success_func func,char* data)
{
    func(data);
}
*/
import "C"
