#include "listeners_bridge.h"
#include "iostream"
// #include "listeners.hpp"

// void *NewOnConnListener() { return new OnConnListener(); }
void OnConnListenerOnConnecting() {

  // Currently connecting to the server. Suitable for displaying a "Connecting"
  // status on the UI.
  std::cout << "On connecting" << std::endl;
}

void OnConnListenerOnConnectSuccess() {
  // Successfully connected to the server.
  std::cout << "On OnConnectSuccess" << std::endl;
}

void OnConnListenerOnConnectFailed(int errCode, const char *errMsg) {
  // Connection to the server failed. You can prompt the user that the current
  // network connection is unavailable.
  std::cout << "Use param.." << errCode << errMsg << std::endl;
  std::cout << "On OnConnectSuccess" << std::endl;
}

void OnConnListenerOnKickedOffline() {
  // The current user has been kicked offline. You can prompt the user with a UI
  // message like "You have logged in to your account on another device. Do you
  // want to log in again?"
  std::cout << "On OnConnectSuccess" << std::endl;
}

void OnConnListenerOnUserTokenExpired() {
  // The login token has expired. Please use a newly issued UserSig for login.
  std::cout << "On OnConnectSuccess" << std::endl;
}