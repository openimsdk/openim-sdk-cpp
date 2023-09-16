// gcc -o test.exe -lc_wrapper.dll test.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#include "openimsdk.h"

typedef struct
{
    GoUint32 platformID;
    char apiAddr[256];
    char wsAddr[256];
    char dataDir[256];
    GoUint32 logLevel;
    GoUint8 isLogStandardOutput;
    char logFilePath[256];
    GoUint8 isExternalExtensions;
} IMConfigC;

void c_conn_callback(int event, char *data)
{
  printf("C c_conn_callback receive from Go callbck code: %d,data: %s\n", event,data);

}
void c_base_callback(char * operationID ,int errCode,char * errMsg,char *data)
{
  printf("C c_base_callback operationID: %s receive from Go callbck code: %d, errMsg: %s, data: %s\n", operationID, errCode, errMsg, data);
}
int main(int argc, char **argv)
{
    char operationID[] = "12345";
//    char uid[] = "6959062403";
    char uid[] = "4365007572";
    //    char token[] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiI2OTU5MDYyNDAzIiwiUGxhdGZvcm1JRCI6MywiZXhwIjoxNzAwNzIwOTg0LCJuYmYiOjE2OTI5NDQ2ODQsImlhdCI6MTY5Mjk0NDk4NH0.8otKTFrOCs8_ueV10rNOD-rzHrCT_EN0obKS9q79bIc";
    char token[] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiI0MzY1MDA3NTcyIiwiUGxhdGZvcm1JRCI6MywiZXhwIjoxNzAxNzYzODQ1LCJuYmYiOjE2OTM5ODc1NDUsImlhdCI6MTY5Mzk4Nzg0NX0.wK4f-R_iuyX7i6gelO2_smob50bdR_2s6FnMvb3xA6U";

    char *jsonString = "{\"platformID\": 3, \"apiAddr\": \"http://125.124.195.201:10002\", \"wsAddr\":\"ws://125.124.195.201:10001\",\"dataDir\": \"./\", \"logLevel\": 1, \"isLogStandardOutput\": true, \"logFilePath\": \"./\", \"isExternalExtensions\": true}";

    GoUint8 init_result;
    init_result = init_sdk(c_conn_callback,operationID, jsonString);
    printf("init_result: %u\n", init_result);

    login(c_base_callback, operationID, uid, token);
    sleep(10);
//    char text[] = "哈哈";
    char* loginUserID=get_login_user();

        printf("return :%s\n",loginUserID);
    char operationID1[] = "12345,create";
    char *message = create_text_message(operationID1, "哈哈");
    printf("return :%s\n",message);
    char operationID2[] = "12345,get_all_conversation_list";
    get_all_conversation_list(c_base_callback, operationID2);

    sleep(1000000);
    return 0;
}