#pragma once
#include <string>

class Base {
public:
  void OnError(int32_t errCode, const std::string &errMsg);
  void OnSuccess(const std::string &data);
};

class SendMsgCallBack : public Base {
public:
  void OnProgress(int progress);
};

class OnConnListener {
public:
  void OnConnecting();
  void OnConnectSuccess();
  void OnConnectFailed(int32_t errCode, const std::string &errMsg);
  void OnKickedOffline();
  void OnUserTokenExpired();
};

// class OnGroupListener {
// public:
//   void OnJoinedGroupAdded(const std::string &groupInfo);
//   void OnJoinedGroupDeleted(const std::string &groupInfo);
//   void OnGroupMemberAdded(const std::string &groupMemberInfo);
//   void OnGroupMemberDeleted(const std::string &groupMemberInfo);
//   void OnGroupApplicationAdded(const std::string &groupApplication);
//   void OnGroupApplicationDeleted(const std::string &groupApplication);
//   void OnGroupInfoChanged(const std::string &groupInfo);
//   void OnGroupMemberInfoChanged(const std::string &groupMemberInfo);
//   void OnGroupApplicationAccepted(const std::string &groupApplication);
//   void OnGroupApplicationRejected(const std::string &groupApplication);
// };

// class OnFriendshipListener {
// public:
//   void OnFriendApplicationAdded(const std::string &friendApplication);
//   void OnFriendApplicationDeleted(const std::string &friendApplication);
//   void OnFriendApplicationAccepted(const std::string &groupApplication);
//   void OnFriendApplicationRejected(const std::string &friendApplication);
//   void OnFriendAdded(const std::string &friendInfo);
//   void OnFriendDeleted(const std::string &friendInfo);
//   void OnFriendInfoChanged(const std::string &friendInfo);
//   void OnBlackAdded(const std::string &blackInfo);
//   void OnBlackDeleted(const std::string &blackInfo);
// };

// class OnConversationListener {
// public:
//   void OnSyncServerStart();
//   void OnSyncServerFinish();
//   void OnSyncServerFailed();
//   void OnNewConversation(const std::string &conversationList);
//   void OnConversationChanged(const std::string &conversationList);
//   void OnTotalUnreadMessageCountChanged(int32_t totalUnreadCount);
// };

// class OnAdvancedMsgListener {
// public:
//   void OnRecvNewMessage(const std::string &message);
//   void OnRecvC2CReadReceipt(const std::string &msgReceiptList);
//   void OnRecvGroupReadReceipt(const std::string &groupMsgReceiptList);
//   void OnRecvMessageRevoked(const std::string &msgID);
//   void OnNewRecvMessageRevoked(const std::string &messageRevoked);
//   void OnRecvMessageExtensionsChanged(const std::string &msgID,
//                                       const std::string
//                                       &reactionExtensionList);
//   void
//   OnRecvMessageExtensionsDeleted(const std::string &msgID,
//                                  const std::string
//                                  &reactionExtensionKeyList);
//   void OnRecvMessageExtensionsAdded(const std::string &msgID,
//                                     const std::string
//                                     &reactionExtensionList);
// };

// class OnBatchMsgListener {
// public:
//   void OnRecvNewMessages(const std::string &messageList);
// };

// class OnUserListener {
// public:
//   void OnSelfInfoUpdated(const std::string &userInfo);
// };

// class OnOrganizationListener {
// public:
//   void OnOrganizationUpdated();
// };

// class OnWorkMomentsListener {
// public:
//   void OnRecvNewNotification();
// };

// class OnCustomBusinessListener {
// public:
//   void OnRecvCustomBusinessMessage(const std::string &businessMessage);
// };

// class OnMessageKvInfoListener {
// public:
//   void OnMessageKvInfoChanged(const std::string &messageChangedList);
// };
