#include "openimsdkcc.h"


namespace openim {

// // ===================================================== init ===============================================
// must be called before use sdk

// instance pattern
OpenIMManager& OpenIMManager::GetInstance()
{
  static OpenIMManager instance;
  return instance;
}

// must be called before use sdk
bool OpenIMManager::InitSDK(const std::function<void(int,const std::string&)>& cCallback, const std::string &operationID, const std::string &config)
{
  char *operationID_cs = const_cast<char *>(operationID.c_str());
  char *config_cs = const_cast<char *>(config.c_str());
  int  ret=init_sdk(_wrapper_cpp_function(cCallback), operationID_cs, config_cs);
  return (ret&1)==1;
}



// release resouces used by SDK
void OpenIMManager::UnInitSDK(const std::string &operationID)
{
  // TODO: free all functions in function pool

  char *operationID_cs = const_cast<char *>(operationID.c_str());
  return un_init_sdk(operationID_cs);
}

// // ===================================================== set listener ===============================================
// impl for set listener, this callback function will be keep in memory,until call SetXXXListener again

// for debug
// void OpenIMManager::SetPrint(const std::function<void(const std::string&)>& printCallBack)
// {
//   this->printCallBack = _wrapper_cpp_function(printCallBack);
//   set_print((CB_S)((this->printCallBack).target<void(*)(const std::string&)>()));
// }

void OpenIMManager::SetAdvancedMsgListener(const std::function<void(int, const std::string &)>& advancedMsgListenerCallback)
{
  set_advanced_msg_listener(_wrapper_cpp_function(advancedMsgListenerCallback));
}
void OpenIMManager::SetBatchMsgListener(const std::function<void(int, const std::string &)>& batchMsgListenerCallback)
{
  set_batch_msg_listener(_wrapper_cpp_function(batchMsgListenerCallback));
}
void OpenIMManager::SetConversationListener(const std::function<void(int, const std::string &)>& conversationListenerCallback)
{
  set_conversation_listener(_wrapper_cpp_function(conversationListenerCallback));
}
void OpenIMManager::SetCustomBusinessListener(const std::function<void(int, const std::string &)>& customBusinessListenerCallback)
{
 set_custom_business_listener(_wrapper_cpp_function(customBusinessListenerCallback));
}
void OpenIMManager::SetFriendListener(const std::function<void(int, const std::string &)>& friendListenerCallback)
{
  set_friend_listener(_wrapper_cpp_function(friendListenerCallback));
}
void OpenIMManager::SetGroupListener(const std::function<void(int, const std::string &)>& groupListenerCallback)
{
  set_group_listener(_wrapper_cpp_function(groupListenerCallback));
}
void OpenIMManager::SetUserListener(const std::function<void(int, const std::string &)>& userListenerCallback)
{
  set_user_listener(_wrapper_cpp_function(userListenerCallback));
}

// // ===================================================== CallOnce Callback ===============================================
// callback function arg below will be free after call once, so we need to wrapp it to onetime
// // ===================================================== message ===============================================

// // ===================================================== login logout ===============================================
void OpenIMManager::Login(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& loginCallback, const std::string& operationID, const std::string& uid, const std::string& token)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* uid_cs=const_cast<char*>(uid.c_str());
  char* token_cs=const_cast<char*>(token.c_str());
  login(_wrapper_callonce_cpp_function(loginCallback),operationID_cs,uid_cs,token_cs);
}

void OpenIMManager::Logout(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& logoutCallback, const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  logout(_wrapper_callonce_cpp_function(logoutCallback),operationID_cs);
}

GoInt OpenIMManager::GetLoginStatus(const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  return get_login_status(operationID_cs);
}

std::string OpenIMManager::GetLoginUser()
{
  char* user=get_login_user();
  std::string user_str(user);
  free(user);
  return get_login_user();
}

void OpenIMManager::NetworkStatusChanged(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& networkStatusCallback, const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  network_status_changed(_wrapper_callonce_cpp_function(networkStatusCallback),operationID_cs);
}


// // ===================================================== message ===============================================
// //

// create text message
std::string OpenIMManager::CreateTextMessage(const std::string& operationID, const std::string& text)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* text_cs=const_cast<char*>(text.c_str());
  char* result_cs=create_text_message(operationID_cs,text_cs);
  std::string result(result_cs);
  // release dynamic c string memory
  free(result_cs);
  return result;
}

// create advanced text message
std::string OpenIMManager::CreateAdvancedTextMessage(const std::string& operationID, const std::string& text, const std::string& messageEntityList)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* text_cs=const_cast<char*>(text.c_str());
  char* messageEntityList_cs=const_cast<char*>(messageEntityList.c_str());
  char* result_cs=create_advanced_text_message(operationID_cs,text_cs,messageEntityList_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create text at message
std::string OpenIMManager::CreateTextAtMessage(const std::string& operationID, const std::string& text, const std::string& atUserList, const std::string& atUsersInfo, const std::string& message)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* text_cs=const_cast<char*>(text.c_str());
  char* atUserList_cs=const_cast<char*>(atUserList.c_str());
  char* atUsersInfo_cs=const_cast<char*>(atUsersInfo.c_str());
  char* message_cs=const_cast<char*>(message.c_str());
  char* result_cs=create_text_at_message(operationID_cs,text_cs,atUserList_cs,atUsersInfo_cs,message_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create location message
std::string OpenIMManager::CreateLocationMessage(const std::string& operationID, const std::string& description, double longitude, double latitude)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* description_cs=const_cast<char*>(description.c_str());
  char* result_cs=create_location_message(operationID_cs,description_cs,longitude,latitude);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create custom message
std::string OpenIMManager::CreateCustomMessage(const std::string& operationID, const std::string& data, const std::string& extension, const std::string& description)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* data_cs=const_cast<char*>(data.c_str());
  char* extension_cs=const_cast<char*>(extension.c_str());
  char* description_cs=const_cast<char*>(description.c_str());
  char* result_cs=create_custom_message(operationID_cs,data_cs,extension_cs,description_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create quote message
std::string OpenIMManager::CreateQuoteMessage(const std::string& operationID, const std::string& text, const std::string& message)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* text_cs=const_cast<char*>(text.c_str());
  char* message_cs=const_cast<char*>(message.c_str());
  char* result_cs=create_quote_message(operationID_cs,text_cs,message_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;

}

// create advanced quote message
std::string OpenIMManager::CreateAdvancedQuoteMessage(const std::string& operationID, const std::string& text, const std::string& message, const std::string& messageEntityList)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* text_cs=const_cast<char*>(text.c_str());
  char* message_cs=const_cast<char*>(message.c_str());
  char* messageEntityList_cs=const_cast<char*>(messageEntityList.c_str());
  char* result_cs=create_advanced_quote_message(operationID_cs,text_cs,message_cs,messageEntityList_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create card message
std::string OpenIMManager::CreateCardMessage(const std::string& operationID, const std::string& cardInfo)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* cardInfo_cs=const_cast<char*>(cardInfo.c_str());
  char* result_cs=create_card_message(operationID_cs,cardInfo_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create video message from full path
std::string OpenIMManager::CreateVideoMessageFromFullPath(const std::string& operationID, const std::string& videoFullPath, const std::string& videoType, long long int duration, const std::string& snapshotFullPath)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* videoFullPath_cs=const_cast<char*>(videoFullPath.c_str());
  char* videoType_cs=const_cast<char*>(videoType.c_str());
  char* snapshotFullPath_cs=const_cast<char*>(snapshotFullPath.c_str());
  char* result_cs=create_video_message_from_full_path(operationID_cs,videoFullPath_cs,videoType_cs,duration,snapshotFullPath_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create image message from full path
std::string OpenIMManager::CreateImageMessageFromFullPath(const std::string& operationID, const std::string& imageFullPath)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* imageFullPath_cs=const_cast<char*>(imageFullPath.c_str());
  char* result_cs=create_image_message_from_full_path(operationID_cs,imageFullPath_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create sound message from full path  
std::string OpenIMManager::CreateSoundMessageFromFullPath(const std::string& operationID, const std::string& soundPath, long long int duration)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* soundPath_cs=const_cast<char*>(soundPath.c_str());
  char* result_cs=create_sound_message_from_full_path(operationID_cs,soundPath_cs,duration);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create file message from full path
std::string OpenIMManager::CreateFileMessageFromFullPath(const std::string& operationID, const std::string& fileFullPath, const std::string& fileName)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* fileFullPath_cs=const_cast<char*>(fileFullPath.c_str());
  char* fileName_cs=const_cast<char*>(fileName.c_str());
  char* result_cs=create_file_message_from_full_path(operationID_cs,fileFullPath_cs,fileName_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create image message
std::string OpenIMManager::CreateImageMessage(const std::string& operationID, const std::string& imagePath)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* imagePath_cs=const_cast<char*>(imagePath.c_str());
  char* result_cs=create_image_message(operationID_cs,imagePath_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create image message by URL
std::string OpenIMManager::CreateImageMessageByURL(const std::string& operationID, const std::string& sourcePicture, const std::string& bigPicture, const std::string& snapshotPicture)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* sourcePicture_cs=const_cast<char*>(sourcePicture.c_str());
  char* bigPicture_cs=const_cast<char*>(bigPicture.c_str());
  char* snapshotPicture_cs=const_cast<char*>(snapshotPicture.c_str());
  char* result_cs=create_image_message_by_url(operationID_cs,sourcePicture_cs,bigPicture_cs,snapshotPicture_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create sound message by URL
std::string OpenIMManager::CreateSoundMessageByURL(const std::string& operationID, const std::string& soundBaseInfo)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* soundBaseInfo_cs=const_cast<char*>(soundBaseInfo.c_str());
  char* result_cs=create_sound_message_by_url(operationID_cs,soundBaseInfo_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create sound message
std::string OpenIMManager::CreateSoundMessage(const std::string& operationID, const std::string& soundPath, long long int duration)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* soundPath_cs=const_cast<char*>(soundPath.c_str());
  char* result_cs=create_sound_message(operationID_cs,soundPath_cs,duration);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create video message by URL
std::string OpenIMManager::CreateVideoMessageByURL(const std::string& operationID, const std::string& videoBaseInfo)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* videoBaseInfo_cs=const_cast<char*>(videoBaseInfo.c_str());
  char* result_cs=create_video_message_by_url(operationID_cs,videoBaseInfo_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create video message
std::string OpenIMManager::CreateVideoMessage(const std::string& operationID, const std::string& videoPath, const std::string& videoType, long long int duration, const std::string& snapshotPath)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* videoPath_cs=const_cast<char*>(videoPath.c_str());
  char* videoType_cs=const_cast<char*>(videoType.c_str());
  char* snapshotPath_cs=const_cast<char*>(snapshotPath.c_str());
  char* result_cs=create_video_message(operationID_cs,videoPath_cs,videoType_cs,duration,snapshotPath_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create file message by URL
std::string OpenIMManager::CreateFileMessageByURL(const std::string& operationID, const std::string& fileBaseInfo)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* fileBaseInfo_cs=const_cast<char*>(fileBaseInfo.c_str());
  char* result_cs=create_file_message_by_url(operationID_cs,fileBaseInfo_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create file message
std::string OpenIMManager::CreateFileMessage(const std::string& operationID, const std::string& filePath, const std::string& fileName)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* filePath_cs=const_cast<char*>(filePath.c_str());
  char* fileName_cs=const_cast<char*>(fileName.c_str());
  char* result_cs=create_file_message(operationID_cs,filePath_cs,fileName_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create merger message
std::string OpenIMManager::CreateMergerMessage(const std::string& operationID, const std::string& messageList, const std::string& title, const std::string& summaryList)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* messageList_cs=const_cast<char*>(messageList.c_str());
  char* title_cs=const_cast<char*>(title.c_str());
  char* summaryList_cs=const_cast<char*>(summaryList.c_str());
  char* result_cs=create_merger_message(operationID_cs,messageList_cs,title_cs,summaryList_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create face message
std::string OpenIMManager::CreateFaceMessage(const std::string& operationID, int index, const std::string& data)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* data_cs=const_cast<char*>(data.c_str());
  char* result_cs=create_face_message(operationID_cs,index,data_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// create forward message
std::string OpenIMManager::CreateForwardMessage(const std::string& operationID, const std::string& m)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* m_cs=const_cast<char*>(m.c_str());
  char* result_cs=create_forward_message(operationID_cs,m_cs);
  std::string result(result_cs);
  free(result_cs);
  return result;
}

// get all conversation list
void OpenIMManager::GetAllConversationList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& getAllConversationListCallback, const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  get_all_conversation_list(_wrapper_callonce_cpp_function(getAllConversationListCallback),operationID_cs);
}

// get advanced history message list
void OpenIMManager::GetAdvancedHistoryMessageList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& getAdvancedHistoryCallback , const std::string& operationID, const std::string& getMessageOptions)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* getMessageOptions_cs=const_cast<char*>(getMessageOptions.c_str());
  get_advanced_history_message_list(_wrapper_callonce_cpp_function(getAdvancedHistoryCallback),operationID_cs,getMessageOptions_cs);
}

// send message
void SendMessage(const std::function<void(const std::string&, int, const std::string&, const std::string&,int)>& sendMessageCallback, const std::string& operationID, const std::string& message,const std::string& recvID,const std::string& groupID,const std::string& offlinePushInfo)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* message_cs=const_cast<char*>(message.c_str());
  char* recvID_cs=const_cast<char*>(recvID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* offlinePushInfo_cs=const_cast<char*>(offlinePushInfo.c_str());
  send_message(_wrapper_callonce_cpp_function(sendMessageCallback),operationID_cs,message_cs,recvID_cs,groupID_cs,offlinePushInfo_cs);
}

// // ===================================================== user ===============================================
// //

// get users info
void OpenIMManager::GetUsersInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& userIDs)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* userIDs_cs=const_cast<char*>(userIDs.c_str());
  get_users_info(_wrapper_callonce_cpp_function(callback),operationID_cs,userIDs_cs);
}

// get users info from server
void OpenIMManager::GetUsersInfoFromServer(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID,const std::string& userIDs)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* userIDs_cs=const_cast<char*>(userIDs.c_str());
  get_users_info_from_srv(_wrapper_callonce_cpp_function(callback),operationID_cs,userIDs_cs);
}

// set self info
void OpenIMManager::SetSelfInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& selfInfo)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* selfInfo_cs=const_cast<char*>(selfInfo.c_str());
  set_self_info(_wrapper_callonce_cpp_function(callback),operationID_cs,selfInfo_cs);
}

// get self user info
void OpenIMManager::GetSelfUserInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  get_self_user_info(_wrapper_callonce_cpp_function(callback),operationID_cs);
}

// update message sender info
void OpenIMManager::UpdateMessageSenderInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& nickname,const std::string& faceURL)
{
  //TODO
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* nickname_cs=const_cast<char*>(nickname.c_str());
  char* faceURL_cs=const_cast<char*>(faceURL.c_str());
  update_msg_sender_info(_wrapper_callonce_cpp_function(callback),operationID_cs,nickname_cs,faceURL_cs);
}

// subscribe users status
void OpenIMManager::SubscribeUsersStatus(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& userIDs)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* userIDs_cs=const_cast<char*>(userIDs.c_str());
  subscribe_users_status(_wrapper_callonce_cpp_function(callback),operationID_cs,userIDs_cs);
}

// unsubscribe users status
void OpenIMManager::UnsubscribeUsersStatus(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& userIDs)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* userIDs_cs=const_cast<char*>(userIDs.c_str());
  unsubscribe_users_status(_wrapper_callonce_cpp_function(callback),operationID_cs,userIDs_cs);
}

// get subscribe users status
void OpenIMManager::GetSubscribeUsersStatus(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  get_subscribe_users_status(_wrapper_callonce_cpp_function(callback),operationID_cs);
}

// get user status
void OpenIMManager::GetUserStatus(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& userID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* userID_cs=const_cast<char*>(userID.c_str());
  get_user_status(_wrapper_callonce_cpp_function(callback),operationID_cs,userID_cs);
}

// // ===================================================== friend ===============================================
// //

// get specified friends info
void OpenIMManager::GetSpecifiedFriendsInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& userIDs)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* userIDs_cs=const_cast<char*>(userIDs.c_str());
  get_specified_friends_info(_wrapper_callonce_cpp_function(callback),operationID_cs,userIDs_cs);
}

// get friend list
void OpenIMManager::GetFriendList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  get_friend_list(_wrapper_callonce_cpp_function(callback),operationID_cs);
}

// get friend list page
void OpenIMManager::GetFriendListPage(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, int offset, int count)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  get_friend_list_page(_wrapper_callonce_cpp_function(callback),operationID_cs,offset,count);
}

// search friends
void OpenIMManager::SearchFriends(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& searchParam)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* searchParam_cs=const_cast<char*>(searchParam.c_str());
  search_friends(_wrapper_callonce_cpp_function(callback),operationID_cs,searchParam_cs);
}

// check friend
void OpenIMManager::CheckFriend(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& userID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* userID_cs=const_cast<char*>(userID.c_str());
  check_friend(_wrapper_callonce_cpp_function(callback),operationID_cs,userID_cs);
}

// add friend
void OpenIMManager::AddFriend(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& friendInfo, const std::string& addFriendExtraInfo)
{
  char* friendInfo_cs=const_cast<char*>(friendInfo.c_str());
  char* addFriendExtraInfo_cs=const_cast<char*>(addFriendExtraInfo.c_str());
  add_friend(_wrapper_callonce_cpp_function(callback),friendInfo_cs,addFriendExtraInfo_cs);
}

// set friend remark
void OpenIMManager::SetFriendRemark(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& userID, const std::string& remark)
{
  char* userID_cs=const_cast<char*>(userID.c_str());
  char* remark_cs=const_cast<char*>(remark.c_str());
  set_friend_remark(_wrapper_callonce_cpp_function(callback),userID_cs,remark_cs);
}

// delete friend
void OpenIMManager::DeleteFriend(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& userID, const std::string& deleteFriendExtraInfo)
{
  char* userID_cs=const_cast<char*>(userID.c_str());
  char* deleteFriendExtraInfo_cs=const_cast<char*>(deleteFriendExtraInfo.c_str());
  delete_friend(_wrapper_callonce_cpp_function(callback),userID_cs,deleteFriendExtraInfo_cs);
}

// get friend application list as recipant
void OpenIMManager::GetFriendApplicationListAsRecipant(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  get_friend_application_list_as_recipient(_wrapper_callonce_cpp_function(callback),operationID_cs);
}

// get friend application list as applicant
void OpenIMManager::GetFriendApplicationListAsApplicant(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  get_friend_application_list_as_applicant(_wrapper_callonce_cpp_function(callback),operationID_cs);
}

// accept friend application
void OpenIMManager::AcceptFriendApplication(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& friendApplicationID, const std::string& extraInfo)
{
  char* friendApplicationID_cs=const_cast<char*>(friendApplicationID.c_str());
  char* extraInfo_cs=const_cast<char*>(extraInfo.c_str());
  accept_friend_application(_wrapper_callonce_cpp_function(callback),friendApplicationID_cs,extraInfo_cs);
}

// refuse friend application
void OpenIMManager::RefuseFriendApplication(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& friendApplicationID, const std::string& extraInfo)
{
  char* friendApplicationID_cs=const_cast<char*>(friendApplicationID.c_str());
  char* extraInfo_cs=const_cast<char*>(extraInfo.c_str());
  refuse_friend_application(_wrapper_callonce_cpp_function(callback),friendApplicationID_cs,extraInfo_cs);
}

// add black
void OpenIMManager::AddBlack(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& userIDs)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* userIDs_cs=const_cast<char*>(userIDs.c_str());
  add_black(_wrapper_callonce_cpp_function(callback),operationID_cs,userIDs_cs);
}

// get black list
void OpenIMManager::GetBlackList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  get_black_list(_wrapper_callonce_cpp_function(callback),operationID_cs);
}

// remove black
void OpenIMManager::RemoveBlack(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& userIDs)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* userIDs_cs=const_cast<char*>(userIDs.c_str());
  remove_black(_wrapper_callonce_cpp_function(callback),operationID_cs,userIDs_cs);
}

// // ===================================================== group ===============================================
// // 

// create group
void OpenIMManager::CreateGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupReqInfo)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupReqInfo_cs=const_cast<char*>(groupReqInfo.c_str());
  create_group(_wrapper_callonce_cpp_function(callback),operationID_cs,groupReqInfo_cs);
}

// join group
void OpenIMManager::JoinGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& reqMsg, int joinSource)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* reqMsg_cs=const_cast<char*>(reqMsg.c_str());
  join_group(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,reqMsg_cs,joinSource);
}

// quit group
void OpenIMManager::QuitGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& groupID, const std::string& reqMsg)
{
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* reqMsg_cs=const_cast<char*>(reqMsg.c_str());
  quit_group(_wrapper_callonce_cpp_function(callback),groupID_cs,reqMsg_cs);
}

// dismiss group
void OpenIMManager::DismissGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& groupID, const std::string& reqMsg)
{
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* reqMsg_cs=const_cast<char*>(reqMsg.c_str());
  dismiss_group(_wrapper_callonce_cpp_function(callback),groupID_cs,reqMsg_cs);
}

// change group mute
void OpenIMManager::ChangeGroupMute(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& groupID, bool mute)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  change_group_mute(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,mute);
}

// change group member mute
void OpenIMManager::ChangeGroupMemberMute(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& groupID, const std::string& memberID, int mutedSeconds)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* memberID_cs=const_cast<char*>(memberID.c_str());
  change_group_member_mute(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,memberID_cs,mutedSeconds);
}


// set the role level of a group member
void OpenIMManager::SetGroupMemberRoleLevel(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& groupID, const std::string& memberID, int roleLevel)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* memberID_cs=const_cast<char*>(memberID.c_str());
  set_group_member_role_level(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,memberID_cs,roleLevel);
}

// set the information of a group member
void OpenIMManager::SetGroupMemberInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupMemberInfo)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupMemberInfo_cs=const_cast<char*>(groupMemberInfo.c_str());
  set_group_member_info(_wrapper_callonce_cpp_function(callback),operationID_cs,groupMemberInfo_cs);
}

// get Joined Group List
void OpenIMManager::GetJoinedGroupList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  get_joined_group_list(_wrapper_callonce_cpp_function(callback),operationID_cs);
}

// get specified groups info
void OpenIMManager::GetSpecifiedGroupsInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& groupIDList)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupIDList_cs=const_cast<char*>(groupIDList.c_str());
  get_specified_groups_info(_wrapper_callonce_cpp_function(callback),operationID_cs,groupIDList_cs);
}

// search groups
void OpenIMManager::SearchGroups(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& searchParam)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* searchParam_cs=const_cast<char*>(searchParam.c_str());
  search_groups(_wrapper_callonce_cpp_function(callback),operationID_cs,searchParam_cs);
}

// set group info
void OpenIMManager::SetGroupInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& groupInfo)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupInfo_cs=const_cast<char*>(groupInfo.c_str());
  set_group_info(_wrapper_callonce_cpp_function(callback),operationID_cs,groupInfo_cs);
}

// set group verification
void OpenIMManager::SetGroupVerification(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& groupID, int verification)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  set_group_verification(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,verification);
}

// set group look member info
void OpenIMManager::SetGroupLookMemberInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& groupID, int lookInfo)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  set_group_look_member_info(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,lookInfo);
}

// set group apply member friend
void OpenIMManager::SetGroupApplyMemberFriend(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& groupID, int rule)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  set_group_apply_member_friend(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,rule);
}

// get group member list
void OpenIMManager::GetGroupMemberList(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& groupID,int filter,int offset,int count)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  get_group_member_list(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,filter,offset,count);
}


// get group member owner and admin
void OpenIMManager::GetGroupMemberOwnerAndAdmin(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID,const std::string& groupID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  get_group_member_owner_and_admin(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs);
}

// get group application list
void OpenIMManager::GetGroupMemberListByJoinTimeFilter(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, int offset, int count, long long int joinTimeBegin, long long int joinTimeEnd, const std::string& filteruserIDs)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* filteruserIDs_cs=const_cast<char*>(filteruserIDs.c_str());
  get_group_member_list_by_join_time_filter(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,offset,count,joinTimeBegin,joinTimeEnd,filteruserIDs_cs);
}

// get specified group members info
void OpenIMManager::GetSpecifiedGroupMembersInfo(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID,const std::string& groupID, const std::string& memberIDList)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* memberIDList_cs=const_cast<char*>(memberIDList.c_str());
  get_specified_group_members_info(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,memberIDList_cs);
}

// kick group members
void OpenIMManager::KickGroupMember(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& reason, const std::string& userIDs)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* reason_cs=const_cast<char*>(reason.c_str());
  char* userIDs_cs=const_cast<char*>(userIDs.c_str());
  kick_group_member(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,reason_cs,userIDs_cs);
}


// transfers the ownership of a group
void OpenIMManager::TransferGroupOwner(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& groupID, const std::string& memberID, const std::string& notifyMsg)
{
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* memberID_cs=const_cast<char*>(memberID.c_str());
  char* notifyMsg_cs=const_cast<char*>(notifyMsg.c_str());
  transfer_group_owner(_wrapper_callonce_cpp_function(callback),groupID_cs,memberID_cs,notifyMsg_cs);
}

// invites users to a group
void OpenIMManager::InviteUserToGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& reason, const std::string& userIDs)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* reason_cs=const_cast<char*>(reason.c_str());
  char* userIDs_cs=const_cast<char*>(userIDs.c_str());
  invite_user_to_group(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,reason_cs,userIDs_cs);
}

// retrives the group application list as a recipient
void OpenIMManager::GetGroupApplicationListAsRecipient(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  get_group_application_list_as_recipient(_wrapper_callonce_cpp_function(callback),operationID_cs);
}

// retrives the group application list as an applicant
void OpenIMManager::GetGroupApplicationListAsApplicant(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  get_group_application_list_as_applicant(_wrapper_callonce_cpp_function(callback),operationID_cs);
}

// accept a group application
void OpenIMManager::AcceptGroupApplication(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& fromUserID, const std::string& handleMsg)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* fromUserID_cs=const_cast<char*>(fromUserID.c_str());
  char* handleMsg_cs=const_cast<char*>(handleMsg.c_str());
  accept_group_application(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,fromUserID_cs,handleMsg_cs);
}


// refuses a group application
void OpenIMManager::RefuseGroupApplication(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& fromUserID, const std::string& handleMsg)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* fromUserID_cs=const_cast<char*>(fromUserID.c_str());
  char* handleMsg_cs=const_cast<char*>(handleMsg.c_str());
  refuse_group_application(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,fromUserID_cs,handleMsg_cs);
}

// set group member nickname
void OpenIMManager::SetGroupMemberNickname(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& groupID, const std::string& userID, const std::string& groupMemberNickname)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  char* userID_cs=const_cast<char*>(userID.c_str());
  char* groupMemberNickname_cs=const_cast<char*>(groupMemberNickname.c_str());
  set_group_member_nickname(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs,userID_cs,groupMemberNickname_cs);
}

// search group members
void OpenIMManager::SearchGroupMembers(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback, const std::string& operationID, const std::string& searchParam)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* searchParam_cs=const_cast<char*>(searchParam.c_str());
  search_group_members(_wrapper_callonce_cpp_function(callback),operationID_cs,searchParam_cs);
}

// check if the user has joined a certain group
void OpenIMManager::IsJoinGroup(const std::function<void(const std::string&, int, const std::string&, const std::string&)>& callback,const std::string& operationID, const std::string& groupID)
{
  char* operationID_cs=const_cast<char*>(operationID.c_str());
  char* groupID_cs=const_cast<char*>(groupID.c_str());
  is_join_group(_wrapper_callonce_cpp_function(callback),operationID_cs,groupID_cs);
}


}