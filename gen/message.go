package main

const (
	Msg_Error = iota

	Msg_Connecting
	Msg_ConnectSuccess
	Msg_ConnectFailed
	Msg_KickedOffline
	Msg_UserTokenExpired

	Msg_SyncServerStart
	Msg_SyncServerFinish
	Msg_SyncServerFailed
	Msg_NewConversation
	Msg_ConversationChanged
	Msg_TotalUnreadMessageCountChanged
	Msg_ConversationUserInputStatusChanged

	Msg_Advanced_RecvNewMessage
	Msg_Advanced_RecvC2CReadReceipt
	Msg_Advanced_RecvGroupReadReceipt
	Msg_Advanced_NewRecvMessageRevoked
	Msg_Advanced_RecvMessageExtensionsChanged
	Msg_Advanced_RecvMessageExtensionsDeleted
	Msg_Advanced_RecvMessageExtensionsAdded
	Msg_Advanced_RecvOfflineNewMessage
	Msg_Advanced_MsgDeleted
	Msg_Advanced_RecvOnlineOnlyMessage

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

	Msg_ErrorOrSuc
)

const (
	DataType_Empty = iota
	DataType_Int
	DataType_Bool
	DataType_LocalConversation
	DataType_LocalConversation_List
	DataType_GetConversationRecvMessageOptResp_List
	DataType_FindMessageList
	DataType_GetAdvancedHistoryMessageList
	DataType_MsgStruct
	DataType_SearchLocalMessagesCallback
	DataType_FullUserInfo
	DataType_FullUserInfo_List
	DataType_FullUserInfoWithCache
	DataType_FullUserInfoWithCache_List
	DataType_LocalUser
	DataType_LocalUser_List
	DataType_OnlineStatus
	DataType_OnlineStatus_List
	DataType_SearchFriendItem
	DataType_SearchFriendItem_List
	DataType_UserIDResult
	DataType_UserIDResult_List
	DataType_LocalFriendRequest
	DataType_LocalFriendRequest_List
	DataType_LocalBlack
	DataType_LocalBlack_List
	DataType_GroupInfo
	DataType_LocalGroup
	DataType_LocalGroup_List
	DataType_LocalGroupMember
	DataType_LocalGroupMember_List
	DataType_LocalAdminGroupRequest
	DataType_LocalAdminGroupRequest_List
	DataType_LocalGroupRequest
	DataType_LocalGroupRequest_List
)

type Empty struct {
}
type Error struct {
	OperationID string `json:"operationId"`
	ErrCode     int32  `json:"errCode"`
	ErrMsg      string `json:"errMsg"`
}
type Success struct {
	OperationID string `json:"operationId"`
	Data        string `json:"data"`
	DataType    int    `json:"dataType"`
}
type ErrorOrSuccess struct {
	OperationID string `json:"operationId"`
	ErrCode     int32  `json:"errCode"`
	Data        string `json:"data"`
	DataType    int    `json:"dataType"`
	ErrMsg      string `json:"errMsg"`
}

type Progress struct {
	OperationID string `json:"operationId"`
	Progress    int    `json:"progress"`
}

type MsgIDAndList struct {
	Id   string `json:"msgId"`
	List string `json:"list"`
}
