# openim-sdk-cpp

openim-sdk-cpp is the C language export layer for the openim-sdk-core repository, designed to provide a simple integration method for C/C++ developers. With this SDK, developers can easily integrate OpenIM real-time messaging services into their C/C++ applications, enabling efficient handling of instant messaging, chat management, message pushing, and other features.

## Setup
### 1. Setting Up Go Environment
To get started, you need to set up the Go development environment. You can download and install Go from the [official Go website](https://go.dev/).

After installation, verify that Go is correctly installed by running:
```
go version
```
### 2. Installing Mage
mage is a Go-based build tool used for running various tasks within the project. To install mage, follow these steps:

 1. Run the following command to install the mage tool:
```
# mac or linux
./bootstrap_install_mage.sh
# windows
./bootstrap_install_mage.bat 
```
 2. Once installed, you can list all available mage commands by running:
 ```
mage -l
# like below
mage -l  
Targets:
  build*          BuildAll compiles the project for all platforms.
  buildAndroid    compiles the project for Android.               
  buildIOS        compiles the project for iOS.                   
  buildLinux      compiles the project for Linux.                 
  buildWindows    compiles the project for Windows.               
                                                                  
* default target     
```

### 3. Generating C Dynamic Library
You can generate the required C dynamic library for your project using the mage tool. Here’s an example of how to build the C dynamic library:
```
mage buildWindows
Building for Windows...
go: downloading github.com/openimsdk/openim-sdk-core/v3 v3.8.1
go: downloading github.com/openimsdk/protocol v0.0.72-alpha.24
go: downloading github.com/openimsdk/tools v0.0.50-alpha.14
# github.com/openimsdk/protocol/user                          
```
This command will generate the appropriate C dynamic library file for your platform (e.g., libopenimsdk.so, libopenimsdk.dylib, etc.).
# Using the C SDK in Your Project
After generating the dynamic library, you can link it to your C/C++ project. Below are the general steps for using the SDK:
### 1. Initialize the SDK
First, initialize the SDK in your project:

```
extern __declspec(dllexport) GoUint8 init_sdk(CB_I_S cCallback, char* operationID, char* config);
```
### 2. Set Up Listeners
Set up listeners for various events such as message reception, login status, etc.:
```
extern __declspec(dllexport) void set_group_listener(CB_I_S cCallback);
extern __declspec(dllexport) void set_conversation_listener(CB_I_S cCallback);
extern __declspec(dllexport) void set_advanced_msg_listener(CB_I_S cCallback);
extern __declspec(dllexport) void set_batch_msg_listener(CB_I_S cCallback);
extern __declspec(dllexport) void set_user_listener(CB_I_S cCallback);
extern __declspec(dllexport) void set_friend_listener(CB_I_S cCallback);
extern __declspec(dllexport) void set_custom_business_listener(CB_I_S cCallback);
```
### 3. Login
Use the login function to start the sdk:
```
extern __declspec(dllexport) void login(CB_S_I_S_S cCallback, char* operationID, char* uid, char* token);
```
### 4. Call Other Interfaces
```
extern __declspec(dllexport) char* create_text_message(char* operationID, char* text);
extern __declspec(dllexport) void send_message(CB_S_I_S_S_I cCallback, char* operationID, char* message, char* recvID, char* groupID, char* offlinePushInfo, int isOnlineOnly);

```
Once the login is successful, you can use other SDK methods to interact with OpenIM services, such as sending messages, creating groups, and more.

For detailed usage and API reference, please refer to the documentation in the openim-sdk-core repository.

# License
This project is licensed under the Apache 2.0 License. See the [LICENSE](https://github.com/openimsdk/openim-sdk-cpp/blob/main/LICENSE) file for more details.

