package main

const (
	Msg_Error = iota

	Msg_Connecting
	Msg_ConnectSuccess
	Msg_ConectFailed
	Msg_KickedOffline
	Msg_UserTokenExpired

	Msg_SyncServerStart
	Msg_SyncServerFinish
	Msg_SyncServerFailed
	Msg_NewConversation
	Msg_ConversationChanged
	Msg_TotalUnreadMessageCountChanged

	Msg_Advanced_RecvNewMessage
	Msg_Advanced_RecvC2CReadReceipt
	Msg_Advanced_RecvGroupReadReceipt
	Msg_Advanced_NewRecvMessageRevoked
	Msg_Advanced_RecvMessageExtensionsChanged
	Msg_Advanced_RecvMessageExtensionsDeleted
	Msg_Advanced_RecvMessageExtensionsAdded
	Msg_Advanced_RecvOfflineNewMessage
	Msg_Advanced_MsgDeleted

	Msg_Batch_RecvNewMessages
	Msg_Batch_RecvOfflineNewMessages

	Msg_FriendApplicationAdded
	Msg_FriendApplicationDeleted
	Msg_FriendApplicationAccepted
	Msg_FriendApplicationRejected
	Msg_FriendAdded
	Msg_FriendDeleted
	Msg_FriendInfoChanged
	Msg_BlackAdded
	Msg_BlackDeleted

	Msg_JoinedGroupAdded
	Msg_JoinedGroupDeleted
	Msg_GroupMemberAdded
	Msg_GroupMemberDeleted
	Msg_GroupApplicationAdded
	Msg_GroupApplicationDeleted
	Msg_GroupInfoChanged
	Msg_GroupDismissed
	Msg_GroupMemberInfoChanged
	Msg_GroupApplicationAccepted
	Msg_GroupApplicationRejected

	Msg_RecvCustomBusinessMessage

	Msg_SelfInfoUpdated
	Msg_UserStatusChanged

	Msg_SendMessage_Error
	Msg_SendMessage_Success
	Msg_SendMessage_Progress

	Msg_Base_Error
	Msg_Base_Success
)

type Empty struct {
}

type Error struct {
	ErrCode     int32  `json:"errCode"`
	ErrMsg      string `json:"errMsg"`
	OperationID string `json:"operationId"`
}

type Success struct {
	OperationID string `json:"operationId"`
	Data        string `json:"data"`
}
type Progress struct {
	OperationID string `json:"operationId"`
	Progress    int    `json:"progress"`
}

type MsgIDAndList struct {
	Id   string `json:"msgId"`
	List string `json:"list"`
}
