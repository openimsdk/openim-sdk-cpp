#pragma once
#include <stdlib.h>
#ifdef __cplusplus
extern "C" {
#endif

// void *NewOnConnListener();

void OnConnListenerOnConnecting();
void OnConnListenerOnConnectSuccess();
void OnConnListenerOnConnectFailed(
    int errCode,
    const char *errMsg); // Not sure if I can use these types
void OnConnListenerOnKickedOffline();
void OnConnListenerOnUserTokenExpired();

#ifdef __cplusplus
} // extern "C"
#endif