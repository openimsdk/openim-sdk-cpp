// gcc -o test.exe -lc_wrapper.dll test.c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#include "openIM.h"

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

void on_connecting()
{
    printf("on_connecting\n");
}
void on_connect_success()
{
    printf("on_connect_success\n");
}
void on_kick_offline()
{
    printf("on_kick_offline\n");
}
void on_user_token_expired()
{
    printf("on_user_token_expired\n");
}
void on_connect_failed(int err_code, char *err_msg)
{
    char *message = (char *)err_msg;
    printf("Error code: %d\n", err_code);
    printf("Error message: %s\n", message);
}
void success(char *data)
{
    printf("login success : %s\n", data);
}
int main(int argc, char **argv)
{
    char operationID[] = "12345";
    char uid[] = "6959062403";
    char token[] = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiI2OTU5MDYyNDAzIiwiUGxhdGZvcm1JRCI6MywiZXhwIjoxNzAwNzIwOTg0LCJuYmYiOjE2OTI5NDQ2ODQsImlhdCI6MTY5Mjk0NDk4NH0.8otKTFrOCs8_ueV10rNOD-rzHrCT_EN0obKS9q79bIc";

    char *jsonString = "{\"platformID\": 3, \"apiAddr\": \"http://125.124.195.201:10002\", \"wsAddr\":\"ws://125.124.195.201:10001\",\"dataDir\": \"./\", \"logLevel\": 1, \"isLogStandardOutput\": true, \"logFilePath\": \"./\", \"isExternalExtensions\": true}";

    GoUint8 init_result;
    init_result = init_sdk(on_connecting, on_connect_success, on_kick_offline, on_user_token_expired, on_connect_failed, operationID, jsonString);
    printf("init_result: %u\n", init_result);

    login(success, on_connect_failed, operationID, uid, token);
//    char text[] = "哈哈";
    GoString message = create_text_message(operationID,"哈哈");
    printf("return :%s",message);

    sleep(1000000);
    return 0;
}