#include <iostream>
#include <functional>
#include <thread>
#include <chrono>
#include "openimsdkcc.h"
using namespace std;
using namespace openim;

// simple test
int main(){
  auto sdkMgr = OpenIMManager::GetInstance();
  string operationID="12345";
  string uid= "openIM123";
  string token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJvcGVuSU0xMjMiLCJQbGF0Zm9ybUlEIjo2LCJleHAiOjE3MDkxMjQ2NzYsIm5iZiI6MTcwMTM0ODM3NiwiaWF0IjoxNzAxMzQ4Njc2fQ.EqlV5TlpiElYhUOHCEcSrZOWi9ldrUMR1L4q0blvxs0";
  string jsonString="{\"platformID\": 6, \"apiAddr\": \"http://14.29.168.56:10002\", \"wsAddr\":\"ws://14.29.168.56:10001\",\"dataDir\": \"./\", \"logLevel\": 5, \"isLogStandardOutput\": true, \"logFilePath\": \"./\", \"isExternalExtensions\": true}";
  bool init_result = sdkMgr.InitSDK([](int event,const string& data){
    cout<<"init> "<<"event:"<<event<<" data:"<<data<<endl;
  },operationID, jsonString);
  cout<<"init_result:"<<init_result<<endl;

  sdkMgr.SetConversationListener([](int event,const string& data){
    cout<<"conversation> "<<"event:"<<event<<" data:"<<data<<endl;
  });
  sdkMgr.SetAdvancedMsgListener([](int event,const string& data){
    cout<<"advancedMsg> "<<"event:"<<event<<" data:"<<data<<endl;
  });
  sdkMgr.Login([](const string& operationID ,int errCode,const string& errMsg, const string& data){
    cout<<"login> " <<"operationID: "<<operationID<<" ,errCode: "<<errCode << "errMsg: "<<errMsg<< ", data: "<<data<<endl; 
  },operationID,uid,token);
  std::this_thread::sleep_for(std::chrono::seconds(10));

  // can only execute below function after login
  string loginUserID=sdkMgr.GetLoginUser();
  cout<<"loginUserID:"<<loginUserID<<endl;
  string operationID1="12345,create";
  string messge=sdkMgr.CreateTextMessage(operationID1,"CCM");
  string operationID2="12345,get_all_conversation_list";
  sdkMgr.GetAllConversationList([](const string& operationID,int errCode ,const string& errMsg,const string& data){
    cout<<"getAllConversationList> "<<"data:"<<data<<endl;
  },operationID2);
  // std::this_thread::sleep_for(std::chrono::seconds(1000));
  while(true);
  return 0;
} 