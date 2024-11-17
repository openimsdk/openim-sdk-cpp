package main

import (
	"encoding/json"

	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"
)

const (
	API_None int = iota
	API_InitSDK
	API_UnInitSDK
	API_Login
	API_Logout
	API_SetAppBackgroundStatus
	API_NetworkStatusChanged
	API_GetLoginStatus
	API_GetLoginUserID
	API_CreateTextMessage
	API_CreateAdvancedTextMessage
	API_CreateTextAtMessage
	API_CreateLocationMessage
	API_CreateCustomMessage
	API_CreateQuoteMessage
	API_CreateAdvancedQuoteMessage
	API_CreateCardMessage
	API_CreateVideoMessageFromFullPath
	API_CreateImageMessageFromFullPath
	API_CreateSoundMessageFromFullPath
	API_CreateFileMessageFromFullPath
	API_CreateImageMessage
	API_CreateImageMessageByURL
	API_CreateSoundMessageByURL
	API_CreateSoundMessage
	API_CreateVideoMessageByURL
	API_CreateVideoMessage
	API_CreateFileMessageByURL
	API_CreateFileMessage
	API_CreateMergerMessage
	API_CreateFaceMessage
	API_CreateForwardMessage
	API_GetAllConversationList
	API_GetConversationListSplit
	API_GetOneConversation
	API_GetMultipleConversation
	API_HideConversation
	API_SetConversation
	API_SetConversationDraft
	API_GetTotalUnreadMsgCount
	API_GetAtAllTag
	API_GetConversationIDBySessionType
	API_SendMessage
	API_SendMessageNotOss
	API_FindMessageList
	API_GetAdvancedHistoryMessageList
	API_GetAdvancedHistoryMessageListReverse
	API_RevokeMessage
	API_TypingStatusUpdate
	API_MarkConversationMessageAsRead
	API_DeleteMessageFromLocalStorage
	API_DeleteMessage
	API_HideAllConversations
	API_DeleteAllMsgFromLocalAndSvr
	API_DeleteAllMsgFromLocal
	API_ClearConversationAndDeleteAllMsg
	API_DeleteConversationAndDeleteAllMsg
	API_InsertSingleMessageToLocalStorage
	API_InsertGroupMessageToLocalStorage
	API_SearchLocalMessages
	API_SetMessageLocalEx
	API_GetUsersInfo
	API_SetSelfInfo
	API_GetSelfUserInfo
	API_SubscribeUsersStatus
	API_UnsubscribeUsersStatus
	API_GetSubscribeUsersStatus
	API_GetUserStatus
	API_GetSpecifiedFriendsInfo
	API_GetFriendList
	API_GetFriendListPage
	API_SearchFriends
	API_UpdateFriends
	API_CheckFriend
	API_AddFriend
	API_DeleteFriend
	API_GetFriendApplicationListAsRecipient
	API_GetFriendApplicationListAsApplicant
	API_AcceptFriendApplication
	API_RefuseFriendApplication
	API_AddBlack
	API_GetBlackList
	API_RemoveBlack
	API_CreateGroup
	API_JoinGroup
	API_QuitGroup
	API_DismissGroup
	API_ChangeGroupMute
	API_ChangeGroupMemberMute
	API_SetGroupMemberInfo
	API_GetJoinedGroupList
	API_GetJoinedGroupListPage
	API_GetSpecifiedGroupsInfo
	API_SearchGroups
	API_SetGroupInfo
	API_GetGroupMemberList
	API_GetGroupMemberOwnerAndAdmin
	API_GetGroupMemberListByJoinTimeFilter
	API_GetSpecifiedGroupMembersInfo
	API_KickGroupMember
	API_TransferGroupOwner
	API_InviteUserToGroup
	API_GetGroupApplicationListAsRecipient
	API_GetGroupApplicationListAsApplicant
	API_AcceptGroupApplication
	API_RefuseGroupApplication
	API_SearchGroupMembers
	API_IsJoinGroup
	API_GetUsersInGroup
)

type APIFUNC func(string) (string, error)

var apiFuncMap = make(map[int]APIFUNC)

func init() {
	apiFuncMap[API_InitSDK] = init_sdk
	apiFuncMap[API_UnInitSDK] = un_init_sdk
	apiFuncMap[API_Login] = login
	apiFuncMap[API_Logout] = logout
	apiFuncMap[API_SetAppBackgroundStatus] = set_app_background_status
	apiFuncMap[API_NetworkStatusChanged] = network_status_changed
	apiFuncMap[API_GetLoginStatus] = get_login_status
	apiFuncMap[API_GetLoginUserID] = get_login_user_id
	apiFuncMap[API_CreateTextMessage] = create_text_message
	apiFuncMap[API_CreateAdvancedTextMessage] = create_advanced_text_message
	apiFuncMap[API_CreateTextAtMessage] = create_text_at_message
	apiFuncMap[API_CreateLocationMessage] = create_location_message
	apiFuncMap[API_CreateCustomMessage] = create_custom_message
	apiFuncMap[API_CreateQuoteMessage] = create_quote_message
	apiFuncMap[API_CreateAdvancedQuoteMessage] = create_advanced_quote_message
	apiFuncMap[API_CreateCardMessage] = create_card_message
	apiFuncMap[API_CreateVideoMessageFromFullPath] = create_video_message_from_full_path
	apiFuncMap[API_CreateImageMessageFromFullPath] = create_image_message_from_full_path
	apiFuncMap[API_CreateSoundMessageFromFullPath] = create_sound_message_from_full_path
	apiFuncMap[API_CreateFileMessageFromFullPath] = create_file_message_from_full_path
	apiFuncMap[API_CreateImageMessage] = create_image_message
	apiFuncMap[API_CreateImageMessageByURL] = create_image_message_by_url
	apiFuncMap[API_CreateSoundMessage] = create_sound_message
	apiFuncMap[API_CreateSoundMessageByURL] = create_sound_message_by_url
	apiFuncMap[API_CreateVideoMessageByURL] = create_video_message_by_url
	apiFuncMap[API_CreateVideoMessage] = create_video_message
	apiFuncMap[API_CreateFileMessageByURL] = create_file_message_by_url
	apiFuncMap[API_CreateFileMessage] = create_file_message
	apiFuncMap[API_CreateMergerMessage] = create_merger_message
	apiFuncMap[API_CreateFaceMessage] = create_face_message
	apiFuncMap[API_CreateForwardMessage] = create_forward_message
	apiFuncMap[API_GetAllConversationList] = get_all_conversation_list
	apiFuncMap[API_GetConversationListSplit] = get_conversation_list_split
	apiFuncMap[API_GetOneConversation] = get_one_conversation
	apiFuncMap[API_GetMultipleConversation] = get_multiple_conversation
	apiFuncMap[API_HideConversation] = hide_conversation
	apiFuncMap[API_SetConversation] = set_conversation
	apiFuncMap[API_SetConversationDraft] = set_conversation_draft
	apiFuncMap[API_GetTotalUnreadMsgCount] = get_total_unread_msg_count
	apiFuncMap[API_GetAtAllTag] = get_at_all_tag
	apiFuncMap[API_GetConversationIDBySessionType] = get_conversation_id_by_session_type
	apiFuncMap[API_SendMessage] = send_message
	apiFuncMap[API_SendMessageNotOss] = send_message_not_oss
	apiFuncMap[API_FindMessageList] = find_message_list
	apiFuncMap[API_GetAdvancedHistoryMessageList] = get_advanced_history_message_list
	apiFuncMap[API_GetAdvancedHistoryMessageListReverse] = get_advanced_history_message_list_reverse
	apiFuncMap[API_RevokeMessage] = revoke_message
	apiFuncMap[API_TypingStatusUpdate] = typing_status_update
	apiFuncMap[API_MarkConversationMessageAsRead] = mark_conversation_message_as_read
	apiFuncMap[API_DeleteMessageFromLocalStorage] = delete_message_from_local_storage
	apiFuncMap[API_DeleteMessage] = delete_message
	apiFuncMap[API_HideAllConversations] = hide_all_conversations
	apiFuncMap[API_DeleteAllMsgFromLocalAndSvr] = delete_all_msg_from_local_and_svr
	apiFuncMap[API_DeleteAllMsgFromLocal] = delete_all_msg_from_local
	apiFuncMap[API_ClearConversationAndDeleteAllMsg] = clear_conversation_and_delete_all_msg
	apiFuncMap[API_DeleteConversationAndDeleteAllMsg] = delete_conversation_and_delete_all_msg
	apiFuncMap[API_InsertSingleMessageToLocalStorage] = insert_single_message_to_local_storage
	apiFuncMap[API_InsertGroupMessageToLocalStorage] = insert_group_message_to_local_storage
	apiFuncMap[API_SearchLocalMessages] = search_local_messages
	apiFuncMap[API_SetMessageLocalEx] = set_message_local_ex
	apiFuncMap[API_GetUsersInfo] = get_users_info
	apiFuncMap[API_SetSelfInfo] = set_self_info
	apiFuncMap[API_GetSelfUserInfo] = get_self_user_info
	apiFuncMap[API_SubscribeUsersStatus] = subscribe_users_status
	apiFuncMap[API_UnsubscribeUsersStatus] = unsubscribe_users_status
	apiFuncMap[API_GetSubscribeUsersStatus] = get_subscribe_users_status
	apiFuncMap[API_GetUserStatus] = get_user_status
	apiFuncMap[API_GetSpecifiedFriendsInfo] = get_specified_friends_info
	apiFuncMap[API_GetFriendList] = get_friend_list
	apiFuncMap[API_GetFriendListPage] = get_friend_list_page
	apiFuncMap[API_SearchFriends] = search_friends
	apiFuncMap[API_UpdateFriends] = update_friends
	apiFuncMap[API_CheckFriend] = check_friend
	apiFuncMap[API_AddFriend] = add_friend
	apiFuncMap[API_DeleteFriend] = delete_friend
	apiFuncMap[API_GetFriendApplicationListAsRecipient] = get_friend_application_list_as_recipient
	apiFuncMap[API_GetFriendApplicationListAsApplicant] = get_friend_application_list_as_applicant
	apiFuncMap[API_AcceptFriendApplication] = accept_friend_application
	apiFuncMap[API_RefuseFriendApplication] = refuse_friend_application
	apiFuncMap[API_AddBlack] = add_black
	apiFuncMap[API_GetBlackList] = get_black_list
	apiFuncMap[API_RemoveBlack] = remove_black
	apiFuncMap[API_CreateGroup] = create_group
	apiFuncMap[API_JoinGroup] = join_group
	apiFuncMap[API_QuitGroup] = quit_group
	apiFuncMap[API_DismissGroup] = dismiss_group
	apiFuncMap[API_ChangeGroupMute] = change_group_mute
	apiFuncMap[API_ChangeGroupMemberMute] = change_group_member_mute
	apiFuncMap[API_SetGroupMemberInfo] = set_group_member_info
	apiFuncMap[API_GetJoinedGroupList] = get_joined_group_list
	apiFuncMap[API_GetJoinedGroupListPage] = get_joined_group_list_page
	apiFuncMap[API_GetSpecifiedGroupsInfo] = get_specified_groups_info
	apiFuncMap[API_SearchGroups] = search_groups
	apiFuncMap[API_SetGroupInfo] = set_group_info
	apiFuncMap[API_GetGroupMemberList] = get_group_member_list
	apiFuncMap[API_GetGroupMemberOwnerAndAdmin] = get_group_member_owner_and_admin
	apiFuncMap[API_GetGroupMemberListByJoinTimeFilter] = get_group_member_list_by_join_time_filter
	apiFuncMap[API_GetSpecifiedGroupMembersInfo] = get_specified_group_members_info
	apiFuncMap[API_KickGroupMember] = kick_group_member
	apiFuncMap[API_TransferGroupOwner] = transfer_group_owner
	apiFuncMap[API_InviteUserToGroup] = invite_user_to_group
	apiFuncMap[API_GetGroupApplicationListAsRecipient] = get_group_application_list_as_recipient
	apiFuncMap[API_GetGroupApplicationListAsApplicant] = get_group_application_list_as_applicant
	apiFuncMap[API_AcceptGroupApplication] = accept_group_application
	apiFuncMap[API_RefuseGroupApplication] = refuse_group_application
	apiFuncMap[API_SearchGroupMembers] = search_group_members
	apiFuncMap[API_IsJoinGroup] = is_join_group
	apiFuncMap[API_GetUsersInGroup] = get_users_in_group
}

func GetAPIFunc(key int) (APIFUNC, bool) {
	apiFunc, ok := apiFuncMap[key]
	return apiFunc, ok
}

type BaseCallback struct {
	OperationId string
	DataType    int
}

func NewBaseCallback(operationId string, dataType int) *BaseCallback {
	return &BaseCallback{OperationId: operationId, DataType: dataType}
}
func (b BaseCallback) OnError(errCode int32, errMsg string) {
	DispatorMsg(Msg_ActiveCall, ErrorOrSuccess{OperationId: b.OperationId, DataType: b.DataType, Data: "", ErrCode: errCode, ErrMsg: errMsg})
}
func (b BaseCallback) OnSuccess(data string) {
	DispatorMsg(Msg_ActiveCall, ErrorOrSuccess{OperationId: b.OperationId, DataType: b.DataType, Data: data, ErrCode: -1, ErrMsg: ""})
}

// =====================================================init_login===============================================
func init_sdk(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Config      string `json:"config"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	callback := NewConnCallback()
	res := open_im_sdk.InitSDK(callback, args.OperationId, args.Config)
	if res {
		open_im_sdk.SetGroupListener(NewGroupCallback())
		open_im_sdk.SetConversationListener(NewConversationCallback())
		open_im_sdk.SetAdvancedMsgListener(NewAdvancedMsgCallback())
		open_im_sdk.SetBatchMsgListener(NewBatchMessageCallback())
		open_im_sdk.SetUserListener(NewUserCallback())
		open_im_sdk.SetFriendListener(NewFriendCallback())
		open_im_sdk.SetCustomBusinessListener(NewCustomBusinessCallback())
	}
	d, err := json.Marshal(BoolValue{Value: res})
	return string(d), err
}

func un_init_sdk(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	open_im_sdk.UnInitSDK(args.OperationId)
	return "", nil
}

func login(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		UId         string `json:"uid"`
		Token       string `json:"token"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.Login(baseCallback, args.OperationId, args.UId, args.Token)
	return "", nil
}

func logout(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.Logout(baseCallback, args.OperationId)
	return "", nil
}

func set_app_background_status(argsStr string) (string, error) {
	args := struct {
		OperationId  string `json:"operationId"`
		IsBackGround bool   `json:"isBackground"`
	}{}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.SetAppBackgroundStatus(baseCallback, args.OperationId, args.IsBackGround)
	return "", nil
}

func network_status_changed(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.NetworkStatusChanged(baseCallback, args.OperationId)
	return "", nil
}

func get_login_status(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	status := open_im_sdk.GetLoginStatus(args.OperationId)
	d, err := json.Marshal(IntValue{Value: status})
	return string(d), err
}

func get_login_user_id(_ string) (string, error) {
	userId := open_im_sdk.GetLoginUserID()
	jsonStr, err := json.Marshal(StringValue{Value: userId})
	return string(jsonStr), err
}

// =====================================================conversation_msg===============================================

func create_text_message(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Text        string `json:"text"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	message := open_im_sdk.CreateTextMessage(args.OperationId, args.Text)
	return message, nil
}

func create_advanced_text_message(argsStr string) (string, error) {
	args := struct {
		OperationId       string `json:"operationId"`
		Text              string `json:"text"`
		MessageEntityList string `json:"messageEntityList"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateAdvancedTextMessage(args.OperationId, args.Text, args.MessageEntityList), nil
}

func create_text_at_message(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Text        string `json:"text"`
		AtUserList  string `json:"atUserList"`
		AtUsersInfo string `json:"atUsersInfo"`
		Message     string `json:"message"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateTextAtMessage(args.OperationId, args.Text, args.AtUserList, args.AtUsersInfo, args.Message), nil
}

func create_location_message(argsStr string) (string, error) {
	args := struct {
		OperationId string  `json:"operationId"`
		Description string  `json:"description"`
		Longitude   float64 `json:"longitude"`
		Latitude    float64 `json:"latitude"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateLocationMessage(args.OperationId, args.Description, args.Longitude, args.Latitude), nil
}

func create_custom_message(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Data        string `json:"data"`
		Extension   string `json:"extension"`
		Description string `json:"description"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateCustomMessage(args.OperationId, args.Data, args.Extension, args.Description), nil
}

func create_quote_message(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Text        string `json:"text"`
		Message     string `json:"message"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateQuoteMessage(args.OperationId, args.Text, args.Message), nil
}

func create_advanced_quote_message(argsStr string) (string, error) {
	args := struct {
		OperationId       string `json:"operationId"`
		Text              string `json:"text"`
		Message           string `json:"message"`
		MessageEntityList string `json:"messageEntityList"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateAdvancedQuoteMessage(args.OperationId, args.Text, args.Message, args.MessageEntityList), nil
}

func create_card_message(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		CardInfo    string `json:"cardInfo"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateCardMessage(args.OperationId, args.CardInfo), nil
}

func create_video_message_from_full_path(argsStr string) (string, error) {
	args := struct {
		OperationId      string `json:"operationId"`
		VideoFullPath    string `json:"videoFullPath"`
		VideoType        string `json:"videoType"`
		Duration         int64  `json:"duration"`
		SnapshotFullPath string `json:"snapshotFullPath"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateVideoMessageFromFullPath(args.OperationId, args.VideoFullPath, args.VideoType, args.Duration, args.SnapshotFullPath), nil
}

func create_image_message_from_full_path(argsStr string) (string, error) {
	args := struct {
		OperationId   string `json:"operationId"`
		ImageFullPath string `json:"imageFullPath"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateImageMessageFromFullPath(args.OperationId, args.ImageFullPath), nil
}

func create_sound_message_from_full_path(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		SoundPath   string `json:"soundPath"`
		Duration    int64  `json:"duration"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateSoundMessageFromFullPath(args.OperationId, args.SoundPath, args.Duration), nil
}

func create_file_message_from_full_path(argsStr string) (string, error) {
	args := struct {
		OperationId  string `json:"operationId"`
		FileFullPath string `json:"fileFullPath"`
		FileName     string `json:"fileName"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateFileMessageFromFullPath(args.OperationId, args.FileFullPath, args.FileName), nil
}

func create_image_message(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		ImagePath   string `json:"imagePath"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateImageMessage(args.OperationId, args.ImagePath), nil
}

func create_image_message_by_url(argsStr string) (string, error) {
	args := struct {
		OperationId     string `json:"operationId"`
		SourcePath      string `json:"sourcePath"`
		SourcePicture   string `json:"sourcePicture"`
		BigPicture      string `json:"bigPicture"`
		SnapshotPicture string `json:"snapshotPicture"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateImageMessageByURL(args.OperationId, args.SourcePath, args.SourcePicture, args.BigPicture, args.SnapshotPicture), nil
}

func create_sound_message_by_url(argsStr string) (string, error) {
	args := struct {
		OperationId   string `json:"operationId"`
		SoundBaseInfo string `json:"soundBaseInfo"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateSoundMessageByURL(args.OperationId, args.SoundBaseInfo), nil
}

func create_sound_message(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		SoundPath   string `json:"soundPath"`
		Duration    int64  `json:"duration"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateSoundMessage(args.OperationId, args.SoundPath, args.Duration), nil
}

func create_video_message_by_url(argsStr string) (string, error) {
	args := struct {
		OperationId   string `json:"operationId"`
		VideoBaseInfo string `json:"videoBaseInfo"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateVideoMessageByURL(args.OperationId, args.VideoBaseInfo), nil
}

func create_video_message(argsStr string) (string, error) {
	args := struct {
		OperationId  string `json:"operationId"`
		VideoPath    string `json:"videoPath"`
		VideoType    string `json:"videoType"`
		Duration     int64  `json:"duration"`
		SnapshotPath string `json:"snapshotPath"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateVideoMessage(args.OperationId, args.VideoPath, args.VideoType, args.Duration, args.SnapshotPath), nil
}

func create_file_message_by_url(argsStr string) (string, error) {
	args := struct {
		OperationId  string `json:"operationId"`
		FileBaseInfo string `json:"fileBaseInfo"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateFileMessageByURL(args.OperationId, args.FileBaseInfo), nil
}

func create_file_message(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		FilePath    string `json:"filePath"`
		FileName    string `json:"fileName"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateFileMessage(args.OperationId, args.FilePath, args.FileName), nil
}

func create_merger_message(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		MessageList string `json:"messageList"`
		Title       string `json:"title"`
		SummaryList string `json:"summaryList"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateMergerMessage(args.OperationId, args.MessageList, args.Title, args.SummaryList), nil
}

func create_face_message(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Index       int    `json:"index"`
		Data        string `json:"data"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateFaceMessage(args.OperationId, args.Index, args.Data), nil
}

func create_forward_message(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Message     string `json:"message"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	return open_im_sdk.CreateForwardMessage(args.OperationId, args.Message), nil
}

func get_all_conversation_list(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Conversation_List)
	open_im_sdk.GetAllConversationList(baseCallback, args.OperationId)
	return "", nil
}

func get_conversation_list_split(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Offset      int    `json:"offset"`
		Count       int    `json:"count"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Conversation_List)
	open_im_sdk.GetConversationListSplit(baseCallback, args.OperationId, args.Offset, args.Count)
	return "", err
}

func get_one_conversation(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		SessionType int32  `json:"sessionType"`
		SourceId    string `json:"sourceId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Conversation)
	open_im_sdk.GetOneConversation(baseCallback, args.OperationId, args.SessionType, args.SourceId)
	return "", nil
}

func get_multiple_conversation(argsStr string) (string, error) {
	args := struct {
		OperationId        string `json:"operationId"`
		ConversationIdList string `json:"conversationIdList"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Conversation_List)
	open_im_sdk.GetMultipleConversation(baseCallback, args.OperationId, args.ConversationIdList)
	return "", nil
}

func hide_conversation(argsStr string) (string, error) {
	args := struct {
		OperationId    string `json:"operationId"`
		ConversationId string `json:"conversationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.HideConversation(baseCallback, args.OperationId, args.ConversationId)
	return "", nil
}

func set_conversation(argsStr string) (string, error) {
	args := struct {
		OperationId    string `json:"operationId"`
		ConversationId string `json:"conversationId"`
		Req            string `json:"req"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.SetConversation(baseCallback, args.OperationId, args.ConversationId, args.Req)
	return "", nil
}

func set_conversation_draft(argsStr string) (string, error) {
	args := struct {
		OperationId    string `json:"operationId"`
		ConversationId string `json:"conversationId"`
		DraftText      string `json:"draftText"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.SetConversationDraft(baseCallback, args.OperationId, args.ConversationId, args.DraftText)
	return "", nil
}

func get_total_unread_msg_count(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Int)
	open_im_sdk.GetTotalUnreadMsgCount(baseCallback, args.OperationId)
	return "", nil
}

func get_at_all_tag(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	tag := open_im_sdk.GetAtAllTag(args.OperationId)
	jsonStr, err := json.Marshal(StringValue{Value: tag})
	return string(jsonStr), err
}

func get_conversation_id_by_session_type(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		SourceId    string `json:"sourceId"`
		SessionType int    `json:"sessionType"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	conversationId := open_im_sdk.GetConversationIDBySessionType(args.OperationId, args.SourceId, args.SessionType)
	jsonStr, err := json.Marshal(StringValue{Value: conversationId})
	return string(jsonStr), err
}

func send_message(argsStr string) (string, error) {
	args := struct {
		OperationId     string `json:"operationId"`
		Message         string `json:"message"`
		RecvId          string `json:"recvId"`
		GroupId         string `json:"groupId"`
		OfflinePushInfo string `json:"offlinePushInfo"`
		IsOnlineOnly    bool   `json:"isOnlineOnly"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	sendMsgCallback := NewSendMessageCallback(args.OperationId)
	open_im_sdk.SendMessage(sendMsgCallback, args.OperationId, args.Message, args.RecvId, args.GroupId, args.OfflinePushInfo, args.IsOnlineOnly)
	return "", nil
}

func send_message_not_oss(argsStr string) (string, error) {
	args := struct {
		OperationId     string `json:"operationId"`
		Message         string `json:"message"`
		RecvId          string `json:"recvId"`
		GroupId         string `json:"groupId"`
		OfflinePushInfo string `json:"offlinePushInfo"`
		IsOnlineOnly    bool   `json:"isOnlineOnly"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	sendMsgCallback := NewSendMessageCallback(args.OperationId)
	open_im_sdk.SendMessageNotOss(sendMsgCallback, args.OperationId, args.Message, args.RecvId, args.GroupId, args.OfflinePushInfo, args.IsOnlineOnly)
	return "", nil
}

func find_message_list(argsStr string) (string, error) {
	args := struct {
		OperationId        string `json:"operationId"`
		FindMessageOptions string `json:"findMessageOptions"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_FindMessageResult)
	open_im_sdk.FindMessageList(baseCallback, args.OperationId, args.FindMessageOptions)
	return "", nil
}

func get_advanced_history_message_list(argsStr string) (string, error) {
	args := struct {
		OperationId       string `json:"operationId"`
		GetMessageOptions string `json:"getMessageOptions"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_AdvancedHistoryMessageResult)
	open_im_sdk.GetAdvancedHistoryMessageList(baseCallback, args.OperationId, args.GetMessageOptions)
	return "", nil
}

func get_advanced_history_message_list_reverse(argsStr string) (string, error) {
	args := struct {
		OperationId       string `json:"operationId"`
		GetMessageOptions string `json:"getMessageOptions"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_AdvancedHistoryMessageResult)
	open_im_sdk.GetAdvancedHistoryMessageListReverse(baseCallback, args.OperationId, args.GetMessageOptions)
	return "", nil
}

func revoke_message(argsStr string) (string, error) {
	args := struct {
		OperationId    string `json:"operationId"`
		ConversationId string `json:"conversationId"`
		ClientMsgId    string `json:"clientMsgId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.RevokeMessage(baseCallback, args.OperationId, args.ConversationId, args.ClientMsgId)
	return "", nil
}

func typing_status_update(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		RecvId      string `json:"recvId"`
		MsgTip      string `json:"msgTip"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.TypingStatusUpdate(baseCallback, args.OperationId, args.RecvId, args.MsgTip)
	return "", nil
}

func mark_conversation_message_as_read(argsStr string) (string, error) {
	args := struct {
		OperationId    string `json:"operationId"`
		ConversationId string `json:"conversationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.MarkConversationMessageAsRead(baseCallback, args.OperationId, args.ConversationId)
	return "", nil
}

func delete_message_from_local_storage(argsStr string) (string, error) {
	args := struct {
		OperationId    string `json:"operationId"`
		ConversationId string `json:"conversationId"`
		ClientMsgId    string `json:"clientMsgId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.DeleteMessageFromLocalStorage(baseCallback, args.OperationId, args.ConversationId, args.ClientMsgId)
	return "", nil
}

func delete_message(argsStr string) (string, error) {
	args := struct {
		OperationId    string `json:"operationId"`
		ConversationId string `json:"conversationId"`
		ClientMsgId    string `json:"clientMsgId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.DeleteMessage(baseCallback, args.OperationId, args.ConversationId, args.ClientMsgId)
	return "", nil
}

func hide_all_conversations(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.HideAllConversations(baseCallback, args.OperationId)
	return "", nil
}

func delete_all_msg_from_local_and_svr(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.DeleteAllMsgFromLocalAndSvr(baseCallback, args.OperationId)
	return "", nil
}

func delete_all_msg_from_local(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.DeleteAllMsgFromLocal(baseCallback, args.OperationId)
	return "", nil
}

func clear_conversation_and_delete_all_msg(argsStr string) (string, error) {
	args := struct {
		OperationId    string `json:"operationId"`
		ConversationId string `json:"conversationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.ClearConversationAndDeleteAllMsg(baseCallback, args.OperationId, args.ConversationId)
	return "", nil
}

func delete_conversation_and_delete_all_msg(argsStr string) (string, error) {
	args := struct {
		OperationId    string `json:"operationId"`
		ConversationId string `json:"conversationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.DeleteConversationAndDeleteAllMsg(baseCallback, args.OperationId, args.ConversationId)
	return "", nil
}

func insert_single_message_to_local_storage(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Message     string `json:"message"`
		RecvId      string `json:"recvId"`
		SendId      string `json:"sendId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Message)
	open_im_sdk.InsertSingleMessageToLocalStorage(baseCallback, args.OperationId, args.Message, args.RecvId, args.SendId)
	return "", nil
}

func insert_group_message_to_local_storage(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Message     string `json:"message"`
		GroupId     string `json:"groupId"`
		SendId      string `json:"sendId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Message)
	open_im_sdk.InsertGroupMessageToLocalStorage(baseCallback, args.OperationId, args.Message, args.GroupId, args.SendId)
	return "", nil
}

func search_local_messages(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		SearchParam string `json:"searchParam"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_SearchMessagesResult)
	open_im_sdk.SearchLocalMessages(baseCallback, args.OperationId, args.SearchParam)
	return "", nil
}

func set_message_local_ex(argsStr string) (string, error) {
	args := struct {
		OperationId    string `json:"operationId"`
		ConversationId string `json:"conversationId"`
		ClientMsgId    string `json:"clientMsgId"`
		LocalEx        string `json:"localEx"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.SetMessageLocalEx(baseCallback, args.OperationId, args.ConversationId, args.ClientMsgId, args.LocalEx)
	return "", nil
}

// =====================================================user===============================================

func get_users_info(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		UserIds     string `json:"userIds"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_PublicUserInfo_List)
	open_im_sdk.GetUsersInfo(baseCallback, args.OperationId, args.UserIds)
	return "", nil
}

func set_self_info(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		UserInfo    string `json:"userInfo"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.SetSelfInfo(baseCallback, args.OperationId, args.UserInfo)
	return "", nil
}

func get_self_user_info(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_UserInfo)
	open_im_sdk.GetSelfUserInfo(baseCallback, args.OperationId)
	return "", nil
}

func subscribe_users_status(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		UserIds     string `json:"userIds"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_OnlineStatus_List)
	open_im_sdk.SubscribeUsersStatus(baseCallback, args.OperationId, args.UserIds)
	return "", nil
}

func unsubscribe_users_status(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		UserIds     string `json:"userIds"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.UnsubscribeUsersStatus(baseCallback, args.OperationId, args.UserIds)
	return "", nil
}

func get_subscribe_users_status(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_OnlineStatus_List)
	open_im_sdk.GetSubscribeUsersStatus(baseCallback, args.OperationId)
	return "", nil
}

func get_user_status(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		UserIds     string `json:"userIds"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_OnlineStatus_List)
	open_im_sdk.GetUserStatus(baseCallback, args.OperationId, args.UserIds)
	return "", nil
}

// =====================================================friend===============================================
func get_specified_friends_info(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		UserIdList  string `json:"userIdList"`
		FilterBlack bool   `json:"filterBlack"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_FriendInfo_List)
	open_im_sdk.GetSpecifiedFriendsInfo(baseCallback, args.OperationId, args.UserIdList, args.FilterBlack)
	return "", nil
}

func get_friend_list(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		FilterBlack bool   `json:"filterBlack"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_FriendInfo_List)
	open_im_sdk.GetFriendList(baseCallback, args.OperationId, args.FilterBlack)
	return "", nil
}

func get_friend_list_page(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Offset      int32  `json:"offset"`
		Count       int32  `json:"count"`
		FilterBlack bool   `json:"filterBlack"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_FriendInfo_List)
	open_im_sdk.GetFriendListPage(baseCallback, args.OperationId, args.Offset, args.Count, args.FilterBlack)
	return "", nil
}

func search_friends(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		SearchParam string `json:"searchParam"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_SearchFriendItem_List)
	open_im_sdk.SearchFriends(baseCallback, args.OperationId, args.SearchParam)
	return "", nil
}

func update_friends(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Req         string `json:"req"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.UpdateFriends(baseCallback, args.OperationId, args.Req)
	return "", nil
}

func check_friend(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		UserIdList  string `json:"userIdList"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_UserIDResult_List)
	open_im_sdk.CheckFriend(baseCallback, args.OperationId, args.UserIdList)
	return "", nil
}

func add_friend(argsStr string) (string, error) {
	args := struct {
		OperationId  string `json:"operationId"`
		UserIdReqMsg string `json:"userIdReqMsg"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.AddFriend(baseCallback, args.OperationId, args.UserIdReqMsg)
	return "", nil
}

func delete_friend(argsStr string) (string, error) {
	args := struct {
		OperationId  string `json:"operationId"`
		FriendUserId string `json:"friendUserId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.DeleteFriend(baseCallback, args.OperationId, args.FriendUserId)
	return "", nil
}

func get_friend_application_list_as_recipient(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_FriendApplicationInfo_List)
	open_im_sdk.GetFriendApplicationListAsRecipient(baseCallback, args.OperationId)
	return "", nil
}

func get_friend_application_list_as_applicant(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_FriendApplicationInfo_List)
	open_im_sdk.GetFriendApplicationListAsApplicant(baseCallback, args.OperationId)
	return "", nil
}

func accept_friend_application(argsStr string) (string, error) {
	args := struct {
		OperationId     string `json:"operationId"`
		UserIdHandleMsg string `json:"userIdHandleMsg"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.AcceptFriendApplication(baseCallback, args.OperationId, args.UserIdHandleMsg)
	return "", nil
}

func refuse_friend_application(argsStr string) (string, error) {
	args := struct {
		OperationId     string `json:"operationId"`
		UserIdHandleMsg string `json:"userIdHandleMsg"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.RefuseFriendApplication(baseCallback, args.OperationId, args.UserIdHandleMsg)
	return "", nil
}

func add_black(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		BlackUserId string `json:"blackUserId"`
		Ex          string `json:"ex"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.AddBlack(baseCallback, args.OperationId, args.BlackUserId, args.Ex)
	return "", nil
}

func get_black_list(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_BlackInfo_List)
	open_im_sdk.GetBlackList(baseCallback, args.OperationId)
	return "", nil
}

func remove_black(argsStr string) (string, error) {
	args := struct {
		OperationId  string `json:"operationId"`
		RemoveUserId string `json:"removeUserId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.RemoveBlack(baseCallback, args.OperationId, args.RemoveUserId)
	return "", nil
}

// =====================================================group===============================================

func create_group(argsStr string) (string, error) {
	args := struct {
		OperationId  string `json:"operationId"`
		GroupReqInfo string `json:"groupReqInfo"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupInfo)
	open_im_sdk.CreateGroup(baseCallback, args.OperationId, args.GroupReqInfo)
	return "", nil
}

func join_group(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
		ReqMsg      string `json:"reqMsg"`
		JoinSource  int32  `json:"joinSource"`
		Ex          string `json:"ex"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.JoinGroup(baseCallback, args.OperationId, args.GroupId, args.ReqMsg, args.JoinSource, args.Ex)
	return "", nil
}

func quit_group(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.QuitGroup(baseCallback, args.OperationId, args.GroupId)
	return "", nil
}

func dismiss_group(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.DismissGroup(baseCallback, args.OperationId, args.GroupId)
	return "", nil
}

func change_group_mute(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
		IsMute      bool   `json:"isMute"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.ChangeGroupMute(baseCallback, args.OperationId, args.GroupId, args.IsMute)
	return "", nil
}

func change_group_member_mute(argsStr string) (string, error) {
	args := struct {
		OperationId  string `json:"operationId"`
		GroupId      string `json:"groupId"`
		UserId       string `json:"userId"`
		MutedSeconds int    `json:"mutedSeconds"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.ChangeGroupMemberMute(baseCallback, args.OperationId, args.GroupId, args.UserId, args.MutedSeconds)
	return "", nil
}

func set_group_member_info(argsStr string) (string, error) {
	args := struct {
		OperationId     string `json:"operationId"`
		GroupMemberInfo string `json:"groupMemberInfo"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.SetGroupMemberInfo(baseCallback, args.OperationId, args.GroupMemberInfo)
	return "", nil
}

func get_joined_group_list(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupInfo_List)
	open_im_sdk.GetJoinedGroupList(baseCallback, args.OperationId)
	return "", nil
}

func get_joined_group_list_page(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		Offset      int32  `json:"offset"`
		Count       int32  `json:"count"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupInfo_List)
	open_im_sdk.GetJoinedGroupListPage(baseCallback, args.OperationId, args.Offset, args.Count)
	return "", nil
}

func get_specified_groups_info(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupIdList string `json:"groupIdList"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupInfo_List)
	open_im_sdk.GetSpecifiedGroupsInfo(baseCallback, args.OperationId, args.GroupIdList)
	return "", nil
}

func search_groups(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		SearchParam string `json:"searchParam"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupInfo_List)
	open_im_sdk.SearchGroups(baseCallback, args.OperationId, args.SearchParam)
	return "", nil
}

func set_group_info(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupInfo   string `json:"groupInfo"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.SetGroupInfo(baseCallback, args.OperationId, args.GroupInfo)
	return "", nil
}

func get_group_member_list(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
		Filter      int32  `json:"filter"`
		Offset      int32  `json:"offset"`
		Count       int32  `json:"count"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupMember_List)
	open_im_sdk.GetGroupMemberList(baseCallback, args.OperationId, args.GroupId, args.Filter, args.Offset, args.Count)
	return "", nil
}

func get_group_member_owner_and_admin(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupMember_List)
	open_im_sdk.GetGroupMemberOwnerAndAdmin(baseCallback, args.OperationId, args.GroupId)
	return "", nil
}

func get_group_member_list_by_join_time_filter(argsStr string) (string, error) {
	args := struct {
		OperationId      string `json:"operationId"`
		GroupId          string `json:"groupId"`
		Offset           int32  `json:"offset"`
		Count            int32  `json:"count"`
		JoinTimeBegin    int64  `json:"joinTimeBegin"`
		JoinTimeEnd      int64  `json:"joinTimeEnd"`
		FilterUserIdList string `json:"filterUserIdList"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupMember_List)
	open_im_sdk.GetGroupMemberListByJoinTimeFilter(baseCallback, args.OperationId, args.GroupId, args.Offset, args.Count, args.JoinTimeBegin, args.JoinTimeEnd, args.FilterUserIdList)
	return "", nil
}

func get_specified_group_members_info(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
		UserIdList  string `json:"userIdList"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupMember_List)
	open_im_sdk.GetSpecifiedGroupMembersInfo(baseCallback, args.OperationId, args.GroupId, args.UserIdList)
	return "", nil
}

func kick_group_member(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
		Reason      string `json:"reason"`
		UserIdList  string `json:"userIdList"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.KickGroupMember(baseCallback, args.OperationId, args.GroupId, args.Reason, args.UserIdList)
	return "", nil
}

func transfer_group_owner(argsStr string) (string, error) {
	args := struct {
		OperationId    string `json:"operationId"`
		GroupId        string `json:"groupId"`
		NewOwnerUserId string `json:"newOwnerUserId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.TransferGroupOwner(baseCallback, args.OperationId, args.GroupId, args.NewOwnerUserId)
	return "", nil
}

func invite_user_to_group(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
		Reason      string `json:"reason"`
		UserIdList  string `json:"userIdList"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.InviteUserToGroup(baseCallback, args.OperationId, args.GroupId, args.Reason, args.UserIdList)
	return "", nil
}

func get_group_application_list_as_recipient(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupApplicationInfo_List)
	open_im_sdk.GetGroupApplicationListAsRecipient(baseCallback, args.OperationId)
	return "", nil
}

func get_group_application_list_as_applicant(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupApplicationInfo_List)
	open_im_sdk.GetGroupApplicationListAsApplicant(baseCallback, args.OperationId)
	return "", nil
}

func accept_group_application(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
		FromUserId  string `json:"fromUserId"`
		HandleMsg   string `json:"handleMsg"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.AcceptGroupApplication(baseCallback, args.OperationId, args.GroupId, args.FromUserId, args.HandleMsg)
	return "", nil
}

func refuse_group_application(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
		FromUserId  string `json:"fromUserId"`
		HandleMsg   string `json:"handleMsg"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Empty)
	open_im_sdk.RefuseGroupApplication(baseCallback, args.OperationId, args.GroupId, args.FromUserId, args.HandleMsg)
	return "", nil
}

func search_group_members(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		SearchParam string `json:"searchParam"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_GroupMember_List)
	open_im_sdk.SearchGroupMembers(baseCallback, args.OperationId, args.SearchParam)
	return "", nil
}

func is_join_group(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_Bool)
	open_im_sdk.IsJoinGroup(baseCallback, args.OperationId, args.GroupId)
	return "", nil
}

func get_users_in_group(argsStr string) (string, error) {
	args := struct {
		OperationId string `json:"operationId"`
		GroupId     string `json:"groupId"`
		UserIdList  string `json:"userIdList"`
	}{}
	err := json.Unmarshal([]byte(argsStr), &args)
	if err != nil {
		return "", err
	}
	baseCallback := NewBaseCallback(args.OperationId, DataType_StringArray)
	open_im_sdk.GetUsersInGroup(baseCallback, args.OperationId, args.GroupId, args.UserIdList)
	return "", nil
}
