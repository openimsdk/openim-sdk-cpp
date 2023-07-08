#include "listeners.hpp"
#include "iostream"

void OnConnListener::OnConnecting() {

  // Currently connecting to the server. Suitable for displaying a "Connecting"
  // status on the UI.
  std::cout << "On connecting" << std::endl;
}

void OnConnListener::OnConnectSuccess() {
  // Successfully connected to the server.
  std::cout << "On OnConnectSuccess" << std::endl;
}

void OnConnListener::OnConnectFailed(int32_t errCode,
                                     const std::string &errMsg) {
  // Connection to the server failed. You can prompt the user that the current
  // network connection is unavailable.
  std::cout << "On OnConnectSuccess" << std::endl;
}

void OnConnListener::OnKickedOffline() {
  // The current user has been kicked offline. You can prompt the user with a UI
  // message like "You have logged in to your account on another device. Do you
  // want to log in again?"
  std::cout << "On OnConnectSuccess" << std::endl;
}

void OnConnListener::OnUserTokenExpired() {
  // The login token has expired. Please use a newly issued UserSig for login.
  std::cout << "On OnConnectSuccess" << std::endl;
}