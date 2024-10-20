package main

type ConnCallback struct {
}

func NewConnCallback() *ConnCallback {
	return &ConnCallback{}
}
func (c ConnCallback) OnConnecting() {
	DispatorMsg(Msg_Connecting, Empty{})
}
func (c ConnCallback) OnConnectSuccess() {
	DispatorMsg(Msg_ConnectSuccess, Empty{})
}
func (c ConnCallback) OnConnectFailed(errCode int32, errMsg string) {
	DispatorMsg(Msg_ConnectFailed, Error{ErrCode: errCode, ErrMsg: errMsg})
}
func (c ConnCallback) OnKickedOffline() {
	DispatorMsg(Msg_KickedOffline, Empty{})
}
func (c ConnCallback) OnUserTokenExpired() {
	DispatorMsg(Msg_UserTokenExpired, Empty{})
}
func (c ConnCallback) OnUserTokenInvalid(errMsg string) {
	DispatorMsg(Msg_UserTokenInvalid, Error{ErrCode: 0, ErrMsg: errMsg})
}

type ConversationCallback struct {
}

func NewConversationCallback() *ConversationCallback {
	return &ConversationCallback{}
}
func (c ConversationCallback) OnSyncServerStart(reinstalled bool) {
	DispatorMsg(Msg_SyncServerStart, struct {
		Reinstalled bool `json:"reinstalled"`
	}{
		Reinstalled: reinstalled,
	})
}
func (c ConversationCallback) OnSyncServerFinish(reinstalled bool) {
	DispatorMsg(Msg_SyncServerFinish, struct {
		Reinstalled bool `json:"reinstalled"`
	}{
		Reinstalled: reinstalled,
	})
}
func (c ConversationCallback) OnSyncServerProgress(progress int) {
	DispatorMsg(Msg_SyncServerProgress, Progress{Progress: progress})
}
func (c ConversationCallback) OnSyncServerFailed(reinstalled bool) {
	DispatorMsg(Msg_SyncServerFailed, struct {
		Reinstalled bool `json:"reinstalled"`
	}{
		Reinstalled: reinstalled,
	})
}
func (c ConversationCallback) OnNewConversation(conversationList string) {
	DispatorMsg(Msg_NewConversation, conversationList)
}
func (c ConversationCallback) OnConversationChanged(conversationList string) {
	DispatorMsg(Msg_ConversationChanged, conversationList)
}
func (c ConversationCallback) OnTotalUnreadMessageCountChanged(totalUnreadCount int32) {
	DispatorMsg(Msg_TotalUnreadMessageCountChanged, totalUnreadCount)
}
func (c ConversationCallback) OnConversationUserInputStatusChanged(change string) {
	DispatorMsg(Msg_ConversationUserInputStatusChanged, change)
}

type AdvancedMsgCallback struct {
}

func (a AdvancedMsgCallback) OnRecvNewMessage(message string) {
	DispatorMsg(Msg_Advanced_RecvNewMessage, message)
}
func (a AdvancedMsgCallback) OnRecvC2CReadReceipt(msgReceiptList string) {
	DispatorMsg(Msg_Advanced_RecvC2CReadReceipt, msgReceiptList)
}
func (a AdvancedMsgCallback) OnRecvGroupReadReceipt(groupMsgReceiptList string) {
	DispatorMsg(Msg_Advanced_RecvGroupReadReceipt, groupMsgReceiptList)
}
func (a AdvancedMsgCallback) OnNewRecvMessageRevoked(messageRevoked string) {
	DispatorMsg(Msg_Advanced_NewRecvMessageRevoked, messageRevoked)
}
func (a AdvancedMsgCallback) OnRecvMessageExtensionsChanged(msgID string, reactionExtensionList string) {
	DispatorMsg(Msg_Advanced_RecvMessageExtensionsChanged, MsgIDAndList{Id: msgID, List: reactionExtensionList})
}
func (a AdvancedMsgCallback) OnRecvMessageExtensionsDeleted(msgID string, reactionExtensionKeyList string) {
	DispatorMsg(Msg_Advanced_RecvMessageExtensionsDeleted, MsgIDAndList{Id: msgID, List: reactionExtensionKeyList})
}
func (a AdvancedMsgCallback) OnRecvMessageExtensionsAdded(msgID string, reactionExtensionList string) {
	DispatorMsg(Msg_Advanced_RecvMessageExtensionsAdded, MsgIDAndList{Id: msgID, List: reactionExtensionList})
}
func (a AdvancedMsgCallback) OnRecvOfflineNewMessage(message string) {
	DispatorMsg(Msg_Advanced_RecvOfflineNewMessage, message)
}
func (a AdvancedMsgCallback) OnMsgDeleted(message string) {
	DispatorMsg(Msg_Advanced_MsgDeleted, message)
}
func (*AdvancedMsgCallback) OnRecvOnlineOnlyMessage(message string) {
	DispatorMsg(Msg_Advanced_RecvOnlineOnlyMessage, message)
}

func NewAdvancedMsgCallback() *AdvancedMsgCallback {
	return &AdvancedMsgCallback{}
}

type BatchMessageCallback struct {
}

func (b BatchMessageCallback) OnRecvNewMessages(messageList string) {
	DispatorMsg(Msg_Batch_RecvNewMessages, messageList)
}
func (b BatchMessageCallback) OnRecvOfflineNewMessages(messageList string) {
	DispatorMsg(Msg_Batch_RecvOfflineNewMessages, messageList)
}

func NewBatchMessageCallback() *BatchMessageCallback {
	return &BatchMessageCallback{}
}

type FriendCallback struct {
}

func (f FriendCallback) OnFriendApplicationAdded(friendApplication string) {
	DispatorMsg(Msg_FriendApplicationAdded, friendApplication)
}
func (f FriendCallback) OnFriendApplicationDeleted(friendApplication string) {
	DispatorMsg(Msg_FriendApplicationDeleted, friendApplication)
}
func (f FriendCallback) OnFriendApplicationAccepted(friendApplication string) {
	DispatorMsg(Msg_FriendApplicationAccepted, friendApplication)
}
func (f FriendCallback) OnFriendApplicationRejected(friendApplication string) {
	DispatorMsg(Msg_FriendApplicationRejected, friendApplication)
}
func (f FriendCallback) OnFriendAdded(friendInfo string) {
	DispatorMsg(Msg_FriendAdded, friendInfo)
}
func (f FriendCallback) OnFriendDeleted(friendInfo string) {
	DispatorMsg(Msg_FriendDeleted, friendInfo)
}
func (f FriendCallback) OnFriendInfoChanged(friendInfo string) {
	DispatorMsg(Msg_FriendInfoChanged, friendInfo)
}
func (f FriendCallback) OnBlackAdded(blackInfo string) {
	DispatorMsg(Msg_BlackAdded, blackInfo)
}
func (f FriendCallback) OnBlackDeleted(blackInfo string) {
	DispatorMsg(Msg_BlackDeleted, blackInfo)
}

func NewFriendCallback() *FriendCallback {
	return &FriendCallback{}
}

type GroupCallback struct {
}

func NewGroupCallback() *GroupCallback {
	return &GroupCallback{}
}
func (g GroupCallback) OnJoinedGroupAdded(groupInfo string) {
	DispatorMsg(Msg_JoinedGroupAdded, groupInfo)
}
func (g GroupCallback) OnJoinedGroupDeleted(groupInfo string) {
	DispatorMsg(Msg_JoinedGroupDeleted, groupInfo)
}
func (g GroupCallback) OnGroupMemberAdded(groupMemberInfo string) {
	DispatorMsg(Msg_GroupMemberAdded, groupMemberInfo)
}
func (g GroupCallback) OnGroupMemberDeleted(groupMemberInfo string) {
	DispatorMsg(Msg_GroupMemberDeleted, groupMemberInfo)
}
func (g GroupCallback) OnGroupApplicationAdded(groupApplication string) {
	DispatorMsg(Msg_GroupApplicationAdded, groupApplication)
}
func (g GroupCallback) OnGroupApplicationDeleted(groupApplication string) {
	DispatorMsg(Msg_GroupApplicationDeleted, groupApplication)
}
func (g GroupCallback) OnGroupInfoChanged(groupInfo string) {
	DispatorMsg(Msg_GroupInfoChanged, groupInfo)
}
func (g GroupCallback) OnGroupDismissed(groupInfo string) {
	DispatorMsg(Msg_GroupDismissed, groupInfo)
}
func (g GroupCallback) OnGroupMemberInfoChanged(groupMemberInfo string) {
	DispatorMsg(Msg_GroupMemberInfoChanged, groupMemberInfo)
}
func (g GroupCallback) OnGroupApplicationAccepted(groupApplication string) {
	DispatorMsg(Msg_GroupApplicationAccepted, groupApplication)
}
func (g GroupCallback) OnGroupApplicationRejected(groupApplication string) {
	DispatorMsg(Msg_GroupApplicationRejected, groupApplication)
}

type CustomBusinessCallback struct {
}

func (c CustomBusinessCallback) OnRecvCustomBusinessMessage(businessMessage string) {
	DispatorMsg(Msg_RecvCustomBusinessMessage, businessMessage)
}
func NewCustomBusinessCallback() *CustomBusinessCallback {
	return &CustomBusinessCallback{}
}

type UserCallback struct {
}

func (u UserCallback) OnSelfInfoUpdated(userInfo string) {
	DispatorMsg(Msg_SelfInfoUpdated, userInfo)
}
func (u UserCallback) OnUserStatusChanged(statusMap string) {
	DispatorMsg(Msg_UserStatusChanged, statusMap)
}
func (u UserCallback) OnUserCommandAdd(userCommand string) {
	DispatorMsg(Msg_UserCommandAdd, userCommand)
}
func (u UserCallback) OnUserCommandDelete(userCommand string) {
	DispatorMsg(Msg_UserCommandDelete, userCommand)
}
func (u UserCallback) OnUserCommandUpdate(userCommand string) {
	DispatorMsg(Msg_UserCommandUpdate, userCommand)
}
func NewUserCallback() *UserCallback {
	return &UserCallback{}
}

type SendMessageCallback struct {
	operationId string
}

func NewSendMessageCallback(operationId string) *SendMessageCallback {
	return &SendMessageCallback{operationId: operationId}
}

func (s SendMessageCallback) OnError(errCode int32, errMsg string) {
	DispatorMsg(Msg_SendMessage_Error, Error{ErrCode: errCode, ErrMsg: errMsg, OperationId: s.operationId})
}

func (s SendMessageCallback) OnSuccess(data string) {
	DispatorMsg(Msg_SendMessage_Success, Success{OperationId: s.operationId, Data: data, DataType: DataType_Message})
}

func (s SendMessageCallback) OnProgress(progress int) {
	DispatorMsg(Msg_SendMessage_Progress, Progress{OperationId: s.operationId, Progress: progress})
}
