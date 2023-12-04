#pragma onces

#include "openimsdk.h"
#include <functional>
#include <string>
#include <memory>
#include <iostream>
#include <bitset>
#include <atomic>
#include <thread>
#include <chrono>
#include <mutex>

#define MAX_NUM_OF_CB_S 10
#define MAX_NUM_OF_CB_I_S 10
#define MAX_NUM_OF_CB_S_I_S_S 10
#define MAX_NUM_OF_CB_S_I_S_S_I 10
#define SLEEP_TIME_FOR_GET_INDEX 100 //ms

// use recursive template to generate enough function pointer array
// and define c type function interface
// and define manager class to manage function pool
namespace {
  CB_S* _fps_cb_s=new CB_S[MAX_NUM_OF_CB_S];
  CB_I_S* _fps_cb_i_s=new CB_I_S[MAX_NUM_OF_CB_I_S];
  CB_S_I_S_S* _fps_cb_s_i_s_s=new CB_S_I_S_S[MAX_NUM_OF_CB_S_I_S_S];
  CB_S_I_S_S_I* _fps_cb_s_i_s_s_i=new CB_S_I_S_S_I[MAX_NUM_OF_CB_S_I_S_S_I];
  // c type func interface call cpp function
  std::function<void(const std::string&)>* _cpp_function_cb_s=new std::function<void(const std::string&)>[MAX_NUM_OF_CB_S];
  std::function<void(int,const std::string&)>* _cpp_function_cb_i_s=new std::function<void(int,const std::string&)>[MAX_NUM_OF_CB_I_S];
  std::function<void(const std::string&,int,const std::string&,const std::string&)>* _cpp_function_cb_s_i_s_s=new std::function<void(const std::string&,int,const std::string&,const std::string&)>[MAX_NUM_OF_CB_S_I_S_S];
  std::function<void(const std::string&,int,const std::string&,const std::string&,int)>* _cpp_function_cb_s_i_s_s_i=new std::function<void(const std::string&,int,const std::string&,const std::string&,int)>[MAX_NUM_OF_CB_S_I_S_S_I];

  template<int N>
  void _generate_cb_s(){
    _fps_cb_s[N]=[](char* c_str){
      _cpp_function_cb_s[N](std::string(c_str));
    };
    _generate_cb_s<N-1>();
  }
  template<>
  void _generate_cb_s<0>(){
    _fps_cb_s[0]=[](char* c_str){
      _cpp_function_cb_s[0](std::string(c_str));
    };
  }

  template<int N>
  void _generate_cb_i_s(){
    _fps_cb_i_s[N]=[](int code,char* c_str){
      _cpp_function_cb_i_s[N](code,std::string(c_str));
    };
    _generate_cb_i_s<N-1>();
  }
  template<>
  void _generate_cb_i_s<0>(){
    _fps_cb_i_s[0]=[](int code,char* c_str){
      _cpp_function_cb_i_s[0](code,std::string(c_str));
    };
  }
  template<int N>
  void _generate_cb_s_i_s_s(){
    _fps_cb_s_i_s_s[N]=[](char* operationID,int code,char* c_str,char* c_str2){
      _cpp_function_cb_s_i_s_s[N](std::string(operationID),code,std::string(c_str),std::string(c_str2));
    };
    _generate_cb_s_i_s_s<N-1>();
  }
  template<>
  void _generate_cb_s_i_s_s<0>(){
    _fps_cb_s_i_s_s[0]=[](char* operationID,int code,char* c_str,char* c_str2){
      _cpp_function_cb_s_i_s_s[0](std::string(operationID),code,std::string(c_str),std::string(c_str2));
    };
  }
  template<int N>
  void _generate_cb_s_i_s_s_i(){
    _fps_cb_s_i_s_s_i[N]=[](char* operationID,int code,char* c_str,char* c_str2,int c_int){
      _cpp_function_cb_s_i_s_s_i[N](std::string(operationID),code,std::string(c_str),std::string(c_str2),c_int);
    };
    _generate_cb_s_i_s_s_i<N-1>();
  }
  template<>
  void _generate_cb_s_i_s_s_i<0>(){
    _fps_cb_s_i_s_s_i[0]=[](char* operationID,int code,char* c_str,char* c_str2,int c_int){
      _cpp_function_cb_s_i_s_s_i[0](std::string(operationID),code,std::string(c_str),std::string(c_str2),c_int);
    };
  }

  // init global function pointer array
  void init(){
    _generate_cb_s<MAX_NUM_OF_CB_S-1>();
    _generate_cb_i_s<MAX_NUM_OF_CB_I_S-1>();
    _generate_cb_s_i_s_s<MAX_NUM_OF_CB_S_I_S_S-1>();
    _generate_cb_s_i_s_s_i<MAX_NUM_OF_CB_S_I_S_S_I-1>();
  }
  // define sigle instance class to manage function pool
  class FuncPoolManager{
    private:
    // define a global bitmap, and support atomic operation, to manage cb_s pool
    std::bitset<MAX_NUM_OF_CB_S> _cb_s_bitmap;
    std::bitset<MAX_NUM_OF_CB_I_S> _cb_i_s_bitmap;
    std::bitset<MAX_NUM_OF_CB_S_I_S_S> _cb_s_i_s_s_bitmap;
    std::bitset<MAX_NUM_OF_CB_S_I_S_S_I> _cb_s_i_s_s_i_bitmap;
    std::mutex _cb_s_mutex;
    std::mutex _cb_i_s_mutex;
    std::mutex _cb_s_i_s_s_mutex;
    std::mutex _cb_s_i_s_s_i_mutex;
    FuncPoolManager(){
      init();
    }
    FuncPoolManager(const FuncPoolManager&){}
    public:
    static FuncPoolManager& get_instance(){
      static FuncPoolManager instance;
      return instance;
    }
    // get a available cb_s function index
    int get_cb_s_index(){
      std::lock_guard<std::mutex> lock(_cb_s_mutex);
      int index=-1;
      for(int i=0;i<_cb_s_bitmap.size();i++){
        if(_cb_s_bitmap[i]==0){
          _cb_s_bitmap[i]=1;
          index=i;
          break;
        }
      }
      return index;
    }
    // get a available cb_i_s function index
    int get_cb_i_s_index(){
      std::lock_guard<std::mutex> lock(_cb_i_s_mutex);
      int index=-1;
      for(int i=0;i<_cb_i_s_bitmap.size();i++){
        if(_cb_i_s_bitmap[i]==0){
          _cb_i_s_bitmap[i]=1;
          index=i;
          break;
        }
      }
      return index;
    }
    // get a available cb_s_i_s_s function index
    int get_cb_s_i_s_s_index(){
      std::lock_guard<std::mutex> lock(_cb_s_i_s_s_mutex);
      int index=-1;
      for(int i=0;i<_cb_s_i_s_s_bitmap.size();i++){
        if(_cb_s_i_s_s_bitmap[i]==0){
          _cb_s_i_s_s_bitmap[i]=1;
          index=i;
          break;
        }
      }
      return index;
    }
    // get a available cb_s_i_s_s_i function index
    int get_cb_s_i_s_s_i_index(){
      std::lock_guard<std::mutex> lock(_cb_s_i_s_s_i_mutex);
      int index=-1;
      for(int i=0;i<_cb_s_i_s_s_i_bitmap.size();i++){
        if(_cb_s_i_s_s_i_bitmap[i]==0){
          _cb_s_i_s_s_i_bitmap[i]=1;
          index=i;
          break;
        }
      }
      return index;
    }
    // release a available cb_s function index
    int release_cb_s_index(int index){
      std::lock_guard<std::mutex> lock(_cb_s_mutex);
      if(index<0||index>=_cb_s_bitmap.size()){
        return -1;
      }
      _cpp_function_cb_s[index]=nullptr;
      _cb_s_bitmap[index]=0;
      return 0;
    }
    // release a available cb_i_s function index
    int release_cb_i_s_index(int index){
      std::lock_guard<std::mutex> lock(_cb_i_s_mutex);
      if(index<0||index>=_cb_i_s_bitmap.size()){
        return -1;
      }
      _cpp_function_cb_i_s[index]=nullptr;
      _cb_i_s_bitmap[index]=0;
      return 0;
    }
    // release a available cb_s_i_s_s function index
    int release_cb_s_i_s_s_index(int index){
      std::lock_guard<std::mutex> lock(_cb_s_i_s_s_mutex);
      if(index<0||index>=_cb_s_i_s_s_bitmap.size()){
        return -1;
      }
      _cpp_function_cb_s_i_s_s[index]=nullptr;
      _cb_s_i_s_s_bitmap[index]=0;
      return 0;
    }
    // release a available cb_s_i_s_s_i function index
    int release_cb_s_i_s_s_i_index(int index){
      std::lock_guard<std::mutex> lock(_cb_s_i_s_s_i_mutex);
      if(index<0||index>=_cb_s_i_s_s_i_bitmap.size()){
        return -1;
      }
      _cpp_function_cb_s_i_s_s_i[index]=nullptr;
      _cb_s_i_s_s_i_bitmap[index]=0;
      return 0;
    }
  };
  FuncPoolManager& instance=FuncPoolManager::get_instance();

  // wrapper persistent function
  // wrapp CB_S,if function pool is full,return nullptr
  CB_S _wrapper_cpp_function(const std::function<void(const std::string&)>& cpp_function) {
    int index=FuncPoolManager::get_instance().get_cb_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_index();
    }
    _cpp_function_cb_s[index]=cpp_function;
    return _fps_cb_s[index];
  }
  // wrapp CB_I_S
  CB_I_S _wrapper_cpp_function(const std::function<void(int,const std::string&)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_i_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_i_s_index();
    }
    _cpp_function_cb_i_s[index]=cpp_function;
    return _fps_cb_i_s[index];
  }
  // wrapp CB_S_I_S_S
  CB_S_I_S_S _wrapper_cpp_function(const std::function<void(const std::string&,int,const std::string&,const std::string&)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_s_i_s_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_i_s_s_index();
    }
    _cpp_function_cb_s_i_s_s[index]=cpp_function;
    return _fps_cb_s_i_s_s[index];
  }
  // wrapp CB_S_I_S_S_I
  CB_S_I_S_S_I _wrapper_cpp_function(const std::function<void(const std::string&,int,const std::string&,const std::string&,int)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_s_i_s_s_i_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_i_s_s_i_index();
    }
    _cpp_function_cb_s_i_s_s_i[index]=cpp_function;
    return _fps_cb_s_i_s_s_i[index];
  }

  // wrapp function to onetime function
  CB_S _wrapper_callonce_cpp_function(const std::function<void(const std::string&)>& cpp_function) {
    int index=FuncPoolManager::get_instance().get_cb_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_index();
    }
    _cpp_function_cb_s[index]=[cpp_function,index](const std::string& str)->void {
      cpp_function(str);
      FuncPoolManager::get_instance().release_cb_s_index(index);
    };
    return _fps_cb_s[index];
  }
  
  CB_I_S _wrapper_callonce_cpp_function(const std::function<void(int,const std::string&)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_i_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_i_s_index();
    }
    _cpp_function_cb_i_s[index]=[cpp_function,index](int code,const std::string& str)->void {
      cpp_function(code,str);
      FuncPoolManager::get_instance().release_cb_i_s_index(index);
    };
    return _fps_cb_i_s[index];
  }
  
  CB_S_I_S_S _wrapper_callonce_cpp_function(const std::function<void(const std::string&,int,const std::string&,const std::string&)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_s_i_s_s_index();
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_i_s_s_index();
    }
    _cpp_function_cb_s_i_s_s[index]=[cpp_function,index](const std::string& operationID,int code,const std::string& str,const std::string& str2)->void {
      cpp_function(operationID,code,str,str2);
      FuncPoolManager::get_instance().release_cb_s_i_s_s_index(index);
    };
    return _fps_cb_s_i_s_s[index];
  }
  
  CB_S_I_S_S_I _wrapper_callonce_cpp_function(const std::function<void(const std::string&,int,const std::string&,const std::string&,int)>& cpp_function)
  {
    int index=FuncPoolManager::get_instance().get_cb_s_i_s_s_i_index();
    // while loop util get a available index
    while(index<0){
      std::this_thread::sleep_for(std::chrono::milliseconds(SLEEP_TIME_FOR_GET_INDEX));
      index=FuncPoolManager::get_instance().get_cb_s_i_s_s_i_index();
    }
    _cpp_function_cb_s_i_s_s_i[index]=[cpp_function,index](const std::string& operationID,int code,const std::string& str,const std::string& str2,int c_int)->void {
      cpp_function(operationID,code,str,str2,c_int);
      FuncPoolManager::get_instance().release_cb_s_i_s_s_i_index(index);
    };
    return _fps_cb_s_i_s_s_i[index];
  }

}


namespace openim{

class OpenIMManager
{
private:
OpenIMManager(){}
public:
  // instance pattern
  static OpenIMManager& GetInstance();

  //must be called before use sdk
  bool InitSDK(const std::function<void(int, const std::string&)>& cCallback,const std::string& operationID,const std::string& config);

  void UnInitSDK(const std::string& operationID);

  // // set print
  // void SetPrint(const std::function<void(const std::string&)>& printCallBack);

  // set group listener
  void SetGroupListener(const std::function<void(int,const std::string&)>& groupListenerCallBack);

  // set conversation listener
  void SetConversationListener(const std::function<void(int, const std::string &)>& conversationListenerCallback);

  // set advanced msg listener
  void SetAdvancedMsgListener(const std::function<void(int, const std::string &)>& advancedMsgListenerCallback);

  // set batch msg listener
  void SetBatchMsgListener(const std::function<void(int, const std::string &)>& batchMsgListenerCallback);

  // set user listener
  void SetUserListener(const std::function<void(int, const std::string &)>& userListenerCallback);

  // set friend listener
  void SetFriendListener(const std::function<void(int, const std::string &)>& friendListenerCallback);

  // set custom business listener
  void SetCustomBusinessListener(const std::function<void(int, const std::string &)>& customBusinessListenerCallback);

  // login
  void Login(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& loginCallback, const std::string& operationID, const std::string& uid, const std::string& token);

  // logout
  void Logout(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& logoutCallback, const std::string& operationID);

  // network status changed
  void NetworkStatusChanged(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& networkStatusCallback, const std::string& operationID);

  // get login status
  GoInt GetLoginStatus(const std::string& operationID);

  // get login user
  std::string GetLoginUser();

  // create text message
  std::string CreateTextMessage(const std::string& operationID, const std::string& text);

  // create advanced text message
  std::string CreateAdvancedTextMessage(const std::string& operationID, const std::string& text, const std::string& messageEntityList);

  // create text at message
  std::string CreateTextAtMessage(const std::string& operationID, const std::string& text, const std::string& atUserList, const std::string& atUsersInfo, const std::string& message);

  // create location message
  std::string CreateLocationMessage(const std::string& operationID, const std::string& description, double longitude, double latitude);

  // create custom message
  std::string CreateCustomMessage(const std::string& operationID, const std::string& data, const std::string& extension, const std::string& description);

  // create quote message
  std::string CreateQuoteMessage(const std::string& operationID, const std::string& text, const std::string& message);

  
  // create advanced quote message
  std::string CreateAdvancedQuoteMessage(const std::string& operationID, const std::string& text, const std::string& message, const std::string& messageEntityList);

  // create card message
  std::string CreateCardMessage(const std::string& operationID, const std::string& cardInfo);

  // create video message from full path
  std::string CreateVideoMessageFromFullPath(const std::string& operationID, const std::string& videoFullPath, const std::string& videoType, long long int duration, const std::string& snapshotFullPath);

  // create image message from full path
  std::string CreateImageMessageFromFullPath(const std::string& operationID, const std::string& imageFullPath);

  // create sound message from full path
  std::string CreateSoundMessageFromFullPath(const std::string& operationID, const std::string& soundPath, long long int duration);

  // create file message from full path
  std::string CreateFileMessageFromFullPath(const std::string& operationID, const std::string& fileFullPath, const std::string& fileName);

  // create image message
  std::string CreateImageMessage(const std::string& operationID, const std::string& imagePath);


  // create image message by URL
  std::string CreateImageMessageByURL(const std::string& operationID, const std::string& sourcePicture, const std::string& bigPicture, const std::string& snapshotPicture);

  // create sound message by URL
  std::string CreateSoundMessageByURL(const std::string& operationID, const std::string& soundBaseInfo);

  // create sound message
  std::string CreateSoundMessage(const std::string& operationID, const std::string& soundPath, long long int duration);

  // create video message by URL
  std::string CreateVideoMessageByURL(const std::string& operationID, const std::string& videoBaseInfo);

  // create video message
  std::string CreateVideoMessage(const std::string& operationID, const std::string& videoPath, const std::string& videoType, long long int duration, const std::string& snapshotPath);

  // create file message by URL
  std::string CreateFileMessageByURL(const std::string& operationID, const std::string& fileBaseInfo);

  // create file message
  std::string CreateFileMessage(const std::string& operationID, const std::string& filePath, const std::string& fileName);

  // create merger message
  std::string CreateMergerMessage(const std::string& operationID, const std::string& messageList, const std::string& title, const std::string& summaryList);

  // create face message
  std::string CreateFaceMessage(const std::string& operationID, int index, const std::string& data);

  // create forward message
  std::string CreateForwardMessage(const std::string& operationID, const std::string& m);

  // get all conversation list
  void GetAllConversationList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& getAllConversationListCallback, const std::string& operationID);
  
  // get advanced history message list
  void GetAdvancedHistoryMessageList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& getAdvancedHistoryCallback , const std::string& operationID, const std::string& getMessageOptions);

  // send message
  void SendMessage(const std::function<void(const std::string&, int, const std::string&, const std::string&, int)>& callback, const std::string& operationID, const std::string& message, const std::string& recvID, const std::string& groupID, const std::string& offlinePushInfo);

  // // =====================================================user===============================================
  // //

  // get users info
  void GetUsersInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDs);

  // get users info from server
  void GetUsersInfoFromServer(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDs);

  // set self info
  void SetSelfInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userInfo);

  // get self user info
  void GetSelfUserInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID);

  // update message sender info
  void UpdateMessageSenderInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& nickname, const std::string& faceURL);

  // subscribe users status
  void SubscribeUsersStatus(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDs);


  // unsubscribe users status
  void UnsubscribeUsersStatus(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDs);

  // get subscribe users status
  void GetSubscribeUsersStatus(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID);
  
  // get user status
  void GetUserStatus(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDs);

  // // =====================================================friend===============================================
  // //

  // get specified friends info
  void GetSpecifiedFriendsInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDs);

  // get friend list
  void GetFriendList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID);

  // get friend list page
  void GetFriendListPage(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, int offset, int count);

  // search friends
  void SearchFriends(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& searchParam);

  // check friend
  void CheckFriend(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDs);

  // add friend
  void AddFriend(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDReqMsg);

  // set friend remark
  void SetFriendRemark(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDRemark);

  // delete friend
  void DeleteFriend(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& friendUserID);

  // get friend application list as recipient
  void GetFriendApplicationListAsRecipant(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID);

  // get friend application list as applicant
  void GetFriendApplicationListAsApplicant(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID);

  // accept friend application
  void AcceptFriendApplication(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDHandleMsg);

  // refuse friend application
  void RefuseFriendApplication(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDHandleMsg);

  // add black
  void AddBlack(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& blackUserID);

  // get black list
  void GetBlackList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID);

  // remove black
  void RemoveBlack(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& removeUserID);


  // // =====================================================group===============================================
  // // 

  // create group
  void CreateGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupReqInfo);

  // join group
  void JoinGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& reqMsg, int joinSource);

  // quit group
  void QuitGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID);

  // dismiss group
  void DismissGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID);

  // change group mute
  void ChangeGroupMute(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, bool isMute);

  // change group member mute
  void ChangeGroupMemberMute(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& userID, int mutedSeconds);

  // SetGroupMemberRoleLevel sets the role level of a group member
  void SetGroupMemberRoleLevel(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& userID, int roleLevel);

  // SetGroupMemberInfo sets the information of a group member
  void SetGroupMemberInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupMemberInfo);

  // GetJoinedGroupList retrieves the list of joined groups
  void GetJoinedGroupList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID);

  // GetSpecifiedGroupsInfo retrieves the information of specified groups
  void GetSpecifiedGroupsInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupIDList);

  // SearchGroups searches for groups
  void SearchGroups(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& searchParam);

  // SetGroupInfo sets the information of a group
  void SetGroupInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupInfo);

  // SetGroupVerification sets the verification mode of a group
  void SetGroupVerification(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, int verification);


  // SetGroupLookMemberInfo sets the member information visibility of a group
  void SetGroupLookMemberInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, int rule);

  // SetGroupApplyMemberFriend sets the friend rule for group applicants
  void SetGroupApplyMemberFriend(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, int rule);

  // GetGroupMemberList retrieves the list of group members
  void GetGroupMemberList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, int filter, int offset, int count);

  // GetGroupMemberOwnerAndAdmin retrieves the owner and admin members of a group
  void GetGroupMemberOwnerAndAdmin(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID);

  // GetGroupMemberListByJoinTimeFilter retrieves the list of group members filtered by join time
  void GetGroupMemberListByJoinTimeFilter(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, int offset, int count, long long int joinTimeBegin, long long int joinTimeEnd, const std::string& filteruserIDs);

  // GetSpecifiedGroupMembersInfo retrieves the information of specified group members
  void GetSpecifiedGroupMembersInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& userIDs);

  // KickGroupMember kicks group members
  void KickGroupMember(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& reason, const std::string& userIDs);

  // TransferGroupOwner transfers the ownership of a group
  void TransferGroupOwner(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& newOwnerUserID);

  // InviteUserToGroup invites users to a group
  void InviteUserToGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& reason, const std::string& userIDs);

  // GetGroupApplicationListAsRecipient retrieves the group application list as a recipient
  void GetGroupApplicationListAsRecipient(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID);

  // GetGroupApplicationListAsApplicant retrieves the group application list as an applicant
  void GetGroupApplicationListAsApplicant(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID);

  // AcceptGroupApplication accepts a group application
  void AcceptGroupApplication(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& fromUserID, const std::string& handleMsg);

  // RefuseGroupApplication refuses a group application
  void RefuseGroupApplication(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& fromUserID, const std::string& handleMsg);

  // SetGroupMemberNickname sets the nickname of a group member
  void SetGroupMemberNickname(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& userID, const std::string& groupMemberNickname);

  // SearchGroupMembers searches for group members
  void SearchGroupMembers(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& searchParam);

  // IsJoinGroup checks if the user has joined a group
  void IsJoinGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID);

};

}