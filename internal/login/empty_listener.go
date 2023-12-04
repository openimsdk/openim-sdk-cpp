package login

type emptyGroupListener struct{}

func newEmptyGroupListener() *emptyGroupListener {
	return &emptyGroupListener{}
}

func (emptyGroupListener) OnJoinedGroupAdded(groupInfo string) {

}

func (emptyGroupListener) OnJoinedGroupDeleted(groupInfo string) {

}

func (emptyGroupListener) OnGroupMemberAdded(groupMemberInfo string) {

}

func (emptyGroupListener) OnGroupMemberDeleted(groupMemberInfo string) {

}

func (emptyGroupListener) OnGroupApplicationAdded(groupApplication string) {

}

func (emptyGroupListener) OnGroupApplicationDeleted(groupApplication string) {

}

func (emptyGroupListener) OnGroupInfoChanged(groupInfo string) {

}

func (emptyGroupListener) OnGroupDismissed(groupInfo string) {

}

func (emptyGroupListener) OnGroupMemberInfoChanged(groupMemberInfo string) {

}

func (emptyGroupListener) OnGroupApplicationAccepted(groupApplication string) {

}

func (emptyGroupListener) OnGroupApplicationRejected(groupApplication string) {

}

type emptyFriendshipListener struct{}

func newEmptyFriendshipListener() *emptyFriendshipListener {
	return &emptyFriendshipListener{}
}

func (emptyFriendshipListener) OnFriendApplicationAdded(friendApplication string) {}

func (emptyFriendshipListener) OnFriendApplicationDeleted(friendApplication string) {}

func (emptyFriendshipListener) OnFriendApplicationAccepted(friendApplication string) {}

func (emptyFriendshipListener) OnFriendApplicationRejected(friendApplication string) {}

func (emptyFriendshipListener) OnFriendAdded(friendInfo string) {}

func (emptyFriendshipListener) OnFriendDeleted(friendInfo string) {}

func (emptyFriendshipListener) OnFriendInfoChanged(friendInfo string) {}

func (emptyFriendshipListener) OnBlackAdded(blackInfo string) {}

func (emptyFriendshipListener) OnBlackDeleted(blackInfo string) {}

type emptyConversationListener struct{}

func newEmptyConversationListener() *emptyConversationListener {
	return &emptyConversationListener{}
}

func (emptyConversationListener) OnSyncServerStart() {

}

func (emptyConversationListener) OnSyncServerFinish() {

}

func (emptyConversationListener) OnSyncServerFailed() {

}

func (emptyConversationListener) OnNewConversation(conversationList string) {

}

func (emptyConversationListener) OnConversationChanged(conversationList string) {

}

func (emptyConversationListener) OnTotalUnreadMessageCountChanged(totalUnreadCount int32) {

}

type emptyAdvancedMsgListener struct{}

func newEmptyAdvancedMsgListener() *emptyAdvancedMsgListener {
	return &emptyAdvancedMsgListener{}
}

func (emptyAdvancedMsgListener) OnRecvNewMessage(message string) {

}

func (emptyAdvancedMsgListener) OnRecvC2CReadReceipt(msgReceiptList string) {

}

func (emptyAdvancedMsgListener) OnRecvGroupReadReceipt(groupMsgReceiptList string) {

}

func (emptyAdvancedMsgListener) OnNewRecvMessageRevoked(messageRevoked string) {

}

func (emptyAdvancedMsgListener) OnRecvMessageExtensionsChanged(msgID string, reactionExtensionList string) {

}

func (emptyAdvancedMsgListener) OnRecvMessageExtensionsDeleted(msgID string, reactionExtensionKeyList string) {

}

func (emptyAdvancedMsgListener) OnRecvMessageExtensionsAdded(msgID string, reactionExtensionList string) {

}

func (emptyAdvancedMsgListener) OnRecvOfflineNewMessage(message string) {

}

func (emptyAdvancedMsgListener) OnMsgDeleted(message string) {

}

type emptyBatchMsgListener struct{}

func newEmptyBatchMsgListener() *emptyBatchMsgListener {
	return &emptyBatchMsgListener{}
}

func (emptyBatchMsgListener) OnRecvNewMessages(messageList string) {

}

func (emptyBatchMsgListener) OnRecvOfflineNewMessages(messageList string) {

}

type emptyUserListener struct{}

func newEmptyUserListener() *emptyUserListener {
	return &emptyUserListener{}
}

func (emptyUserListener) OnSelfInfoUpdated(userInfo string) {

}

func (emptyUserListener) OnUserStatusChanged(statusMap string) {

}

type emptyCustomBusinessListener struct{}

func newEmptyCustomBusinessListener() *emptyCustomBusinessListener {
	return &emptyCustomBusinessListener{}
}

func (emptyCustomBusinessListener) OnRecvCustomBusinessMessage(businessMessage string) {

}
