package main

/*
#include <stdio.h>
typedef void (*MessageHandler)(int id ,char* data);
extern MessageHandler messageHandler;
extern void CallMessageHandler(MessageHandler msgHandler,int id,char* data);
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"
)

//export SetMessageHandler
func SetMessageHandler(handler C.MessageHandler) {
	C.messageHandler = handler
}

func DispatorMsg(msgId int, msg interface{}) {
	t := reflect.TypeOf(msg)
	kind := t.Kind()
	if kind == reflect.Struct {
		msgJson, err := json.Marshal(msg)
		if err != nil {
			C.CallMessageHandler(C.messageHandler, C.int(0), C.CString(fmt.Sprintf("Marshal Json Error :%s", err.Error())))
		} else {
			C.CallMessageHandler(C.messageHandler, C.int(msgId), C.CString((string(msgJson))))
		}
	} else if kind == reflect.String {
		C.CallMessageHandler(C.messageHandler, C.int(msgId), C.CString(msg.(string)))
	} else if kind == reflect.Int32 {
		C.CallMessageHandler(C.messageHandler, C.int(msgId), C.CString(strconv.Itoa(msg.(int))))
	}
}

func parseBool(b int) bool {
	return !(b == 0)
}

// =====================================================listener===============================================

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

type ConversationCallback struct {
}

func NewConversationCallback() *ConversationCallback {
	return &ConversationCallback{}
}
func (c ConversationCallback) OnSyncServerStart() {
	DispatorMsg(Msg_SyncServerStart, Empty{})
}
func (c ConversationCallback) OnSyncServerFinish() {
	DispatorMsg(Msg_SyncServerFinish, Empty{})
}
func (c ConversationCallback) OnSyncServerFailed() {
	DispatorMsg(Msg_SyncServerFailed, Empty{})
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

func NewUserCallback() *UserCallback {
	return &UserCallback{}
}

type SendMessageCallback struct {
	operationID string
}

func NewSendMessageCallback(operationID *C.char) *SendMessageCallback {
	return &SendMessageCallback{operationID: C.GoString(operationID)}
}

func (s SendMessageCallback) OnError(errCode int32, errMsg string) {
	DispatorMsg(Msg_SendMessage_Error, Error{ErrCode: errCode, ErrMsg: errMsg, OperationID: s.operationID})
}

func (s SendMessageCallback) OnSuccess(data string) {
	DispatorMsg(Msg_SendMessage_Success, Success{OperationID: s.operationID, Data: data})
}

func (s SendMessageCallback) OnProgress(progress int) {
	DispatorMsg(Msg_SendMessage_Progress, Progress{OperationID: s.operationID, Progress: progress})
}

type BaseCallback struct {
	OperationID string
	DataType    int
}

func NewBaseCallback(operationID *C.char, dataType int) *BaseCallback {
	return &BaseCallback{OperationID: C.GoString(operationID), DataType: dataType}
}
func (b BaseCallback) OnError(errCode int32, errMsg string) {
	DispatorMsg(Msg_Base_Error, Error{OperationID: b.OperationID, ErrCode: errCode, ErrMsg: errMsg})
}
func (b BaseCallback) OnSuccess(data string) {
	DispatorMsg(Msg_Base_Success, Success{OperationID: b.OperationID, DataType: b.DataType, Data: data})
}

type GroupListener struct {
}

func (GroupListener) OnGroupApplicationAccepted(groupApplication string) {

}
func (GroupListener) OnGroupApplicationAdded(groupApplication string) {
}
func (GroupListener) OnGroupApplicationDeleted(groupApplication string) {
}

func (GroupListener) OnGroupApplicationRejected(groupApplication string) {
}

func (GroupListener) OnGroupDismissed(groupInfo string) {
}

func (GroupListener) OnGroupInfoChanged(groupInfo string) {
}

func (GroupListener) OnGroupMemberAdded(groupMemberInfo string) {
}
func (GroupListener) OnGroupMemberDeleted(groupMemberInfo string) {
}
func (GroupListener) OnGroupMemberInfoChanged(groupMemberInfo string) {
}

func (GroupListener) OnJoinedGroupAdded(groupInfo string) {
}

func (GroupListener) OnJoinedGroupDeleted(groupInfo string) {
}

type ConversationListener struct {
}

func (ConversationListener) OnConversationChanged(conversationList string) {
}
func (ConversationListener) OnNewConversation(conversationList string) {
}
func (ConversationListener) OnSyncServerFailed() {
}
func (ConversationListener) OnSyncServerFinish() {
}
func (ConversationListener) OnSyncServerStart() {
}

func (ConversationListener) OnTotalUnreadMessageCountChanged(totalUnreadCount int32) {
}

type AdvancedMsgListener struct {
}

func (AdvancedMsgListener) OnMsgDeleted(message string) {
}
func (AdvancedMsgListener) OnNewRecvMessageRevoked(messageRevoked string) {
}
func (AdvancedMsgListener) OnRecvC2CReadReceipt(msgReceiptList string) {
}
func (AdvancedMsgListener) OnRecvGroupReadReceipt(groupMsgReceiptList string) {
}
func (AdvancedMsgListener) OnRecvMessageExtensionsAdded(msgID string, reactionExtensionList string) {
}
func (AdvancedMsgListener) OnRecvMessageExtensionsChanged(msgID string, reactionExtensionList string) {
}
func (AdvancedMsgListener) OnRecvMessageExtensionsDeleted(msgID string, reactionExtensionKeyList string) {
}
func (AdvancedMsgListener) OnRecvNewMessage(message string) {
}
func (AdvancedMsgListener) OnRecvOfflineNewMessage(message string) {
}

type BatchMsgListener struct {
}

func (BatchMsgListener) OnRecvNewMessages(messageList string) {
}

func (BatchMsgListener) OnRecvOfflineNewMessages(messageList string) {
}

type UserListener struct {
}

func (UserListener) OnSelfInfoUpdated(userInfo string) {
}

func (UserListener) OnUserStatusChanged(userOnlineStatus string) {

}

type FriendListener struct {
}

func (FriendListener) OnBlackAdded(blackInfo string) {

}

func (FriendListener) OnBlackDeleted(blackInfo string) {

}

func (FriendListener) OnFriendAdded(friendInfo string) {

}

func (FriendListener) OnFriendApplicationAccepted(friendApplication string) {

}

func (FriendListener) OnFriendApplicationAdded(friendApplication string) {

}

func (FriendListener) OnFriendApplicationDeleted(friendApplication string) {

}

func (FriendListener) OnFriendApplicationRejected(friendApplication string) {

}

func (FriendListener) OnFriendDeleted(friendInfo string) {

}

func (FriendListener) OnFriendInfoChanged(friendInfo string) {

}

type CustomBusinessListener struct {
}

func (CustomBusinessListener) OnRecvCustomBusinessMessage(businessMessage string) {

}

type MessageKvInfoListener struct {
}

func (MessageKvInfoListener) OnMessageKvInfoChanged(messageChangedList string) {

}

// =====================================================init_login===============================================

//export  init_sdk
func init_sdk(
	operationID *C.char, config *C.char) bool {
	callback := NewConnCallback()
	res := open_im_sdk.InitSDK(callback, C.GoString(operationID), C.GoString(config))
	if res {
		open_im_sdk.SetGroupListener(NewGroupCallback())
		open_im_sdk.SetConversationListener(NewConversationCallback())
		open_im_sdk.SetAdvancedMsgListener(NewAdvancedMsgCallback())
		open_im_sdk.SetBatchMsgListener(NewBatchMessageCallback())
		open_im_sdk.SetUserListener(NewUserCallback())
		open_im_sdk.SetFriendListener(NewFriendCallback())
		open_im_sdk.SetCustomBusinessListener(NewCustomBusinessCallback())
	}
	return res
}

//export un_init_sdk
func un_init_sdk(operationID *C.char) {
	open_im_sdk.UnInitSDK(C.GoString(operationID))
}

//export  login
func login(operationID, uid, token *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.Login(baseCallback, C.GoString(operationID), C.GoString(uid), C.GoString(token))
}

//export  logout
func logout(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.Logout(baseCallback, C.GoString(operationID))
}

//export set_app_background_status
func set_app_background_status(operationID *C.char, isBackground C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetAppBackgroundStatus(baseCallback, C.GoString(operationID), parseBool(int(isBackground)))
}

//export  network_status_changed
func network_status_changed(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.NetworkStatusChanged(baseCallback, C.GoString(operationID))
}

//export  get_login_status
func get_login_status(operationID *C.char) int {
	return open_im_sdk.GetLoginStatus(C.GoString(operationID))
}

//export  get_login_user
func get_login_user() *C.char {
	return C.CString(open_im_sdk.GetLoginUserID())
}

// =====================================================conversation_msg===============================================

//export  create_text_message
func create_text_message(operationID, text *C.char) *C.char {
	message := C.CString(open_im_sdk.CreateTextMessage(C.GoString(operationID), C.GoString(text)))
	return message
}

//export create_advanced_text_message
func create_advanced_text_message(operationID, text, messageEntityList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateAdvancedTextMessage(C.GoString(operationID), C.GoString(text),
		C.GoString(messageEntityList)))
}

//export create_text_at_message
func create_text_at_message(operationID, text, atUserList, atUsersInfo, message *C.char) *C.char {
	return C.CString(open_im_sdk.CreateTextAtMessage(C.GoString(operationID), C.GoString(text), C.GoString(atUserList),
		C.GoString(atUsersInfo), C.GoString(message)))
}

//export create_location_message
func create_location_message(operationID, description *C.char, longitude, latitude C.double) *C.char {
	return C.CString(open_im_sdk.CreateLocationMessage(C.GoString(operationID), C.GoString(description),
		float64(longitude), float64(latitude)))
}

//export create_custom_message
func create_custom_message(operationID, data, extension, description *C.char) *C.char {
	return C.CString(open_im_sdk.CreateCustomMessage(C.GoString(operationID), C.GoString(data), C.GoString(extension),
		C.GoString(description)))
}

//export create_quote_message
func create_quote_message(operationID, text, message *C.char) *C.char {
	return C.CString(open_im_sdk.CreateQuoteMessage(C.GoString(operationID), C.GoString(text), C.GoString(message)))
}

//export create_advanced_quote_message
func create_advanced_quote_message(operationID, text, message, messageEntityList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateAdvancedQuoteMessage(C.GoString(operationID), C.GoString(text),
		C.GoString(message), C.GoString(messageEntityList)))
}

//export create_card_message
func create_card_message(operationID, cardInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateCardMessage(C.GoString(operationID), C.GoString(cardInfo)))
}

//export create_video_message_from_full_path
func create_video_message_from_full_path(operationID, videoFullPath, videoType *C.char, duration C.longlong,
	snapshotFullPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessageFromFullPath(C.GoString(operationID), C.GoString(videoFullPath),
		C.GoString(videoType), int64(duration), C.GoString(snapshotFullPath)))
}

//export create_image_message_from_full_path
func create_image_message_from_full_path(operationID, imageFullPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessageFromFullPath(C.GoString(operationID), C.GoString(imageFullPath)))
}

//export create_sound_message_from_full_path
func create_sound_message_from_full_path(operationID, soundPath *C.char, duration C.longlong) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessageFromFullPath(C.GoString(operationID), C.GoString(soundPath),
		int64(duration)))
}

//export create_file_message_from_full_path
func create_file_message_from_full_path(operationID, fileFullPath, fileName *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessageFromFullPath(C.GoString(operationID), C.GoString(fileFullPath),
		C.GoString(fileName)))
}

//export create_image_message
func create_image_message(operationID, imagePath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessage(C.GoString(operationID), C.GoString(imagePath)))
}

//export create_image_message_by_url
func create_image_message_by_url(operationID, sourcePath, sourcePicture, bigPicture, snapshotPicture *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessageByURL(C.GoString(operationID), C.GoString(sourcePath),
		C.GoString(sourcePicture), C.GoString(bigPicture), C.GoString(snapshotPicture)))
}

//export create_sound_message_by_url
func create_sound_message_by_url(operationID, soundBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessageByURL(C.GoString(operationID), C.GoString(soundBaseInfo)))
}

//export create_sound_message
func create_sound_message(operationID, soundPath *C.char, duration C.longlong) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessage(C.GoString(operationID), C.GoString(soundPath), int64(duration)))
}

//export create_video_message_by_url
func create_video_message_by_url(operationID, videoBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessageByURL(C.GoString(operationID), C.GoString(videoBaseInfo)))
}

//export create_video_message
func create_video_message(operationID, videoPath *C.char, videoType *C.char, duration C.longlong,
	snapshotPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessage(C.GoString(operationID), C.GoString(videoPath),
		C.GoString(videoType), int64(duration), C.GoString(snapshotPath)))
}

//export create_file_message_by_url
func create_file_message_by_url(operationID, fileBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessageByURL(C.GoString(operationID), C.GoString(fileBaseInfo)))
}

//export create_file_message
func create_file_message(operationID, filePath, fileName *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessage(C.GoString(operationID), C.GoString(filePath), C.GoString(fileName)))
}

//export create_merger_message
func create_merger_message(operationID, messageList, title, summaryList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateMergerMessage(C.GoString(operationID), C.GoString(messageList),
		C.GoString(title), C.GoString(summaryList)))
}

//export create_face_message
func create_face_message(operationID *C.char, index C.int, data *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFaceMessage(C.GoString(operationID), int(index), C.GoString(data)))
}

//export create_forward_message
func create_forward_message(operationID, m *C.char) *C.char {
	return C.CString(open_im_sdk.CreateForwardMessage(C.GoString(operationID), C.GoString(m)))
}

//export get_all_conversation_list
func get_all_conversation_list(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalConversation_List)
	open_im_sdk.GetAllConversationList(baseCallback, C.GoString(operationID))
}

//export get_conversation_list_split
func get_conversation_list_split(operationID *C.char, offset C.int, count C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalConversation_List)
	open_im_sdk.GetConversationListSplit(baseCallback, C.GoString(operationID), int(offset), int(count))
}

//export get_one_conversation
func get_one_conversation(operationID *C.char, sessionType C.int, sourceID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalConversation)
	open_im_sdk.GetOneConversation(baseCallback, C.GoString(operationID), int32(sessionType), C.GoString(sourceID))
}

//export get_multiple_conversation
func get_multiple_conversation(operationID *C.char, conversationIDList *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalConversation_List)
	open_im_sdk.GetMultipleConversation(baseCallback, C.GoString(operationID), C.GoString(conversationIDList))
}

//export set_conversation_msg_destruct_time
func set_conversation_msg_destruct_time(operationID *C.char, conversationID *C.char, msgDestructTime C.longlong) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetConversationMsgDestructTime(baseCallback, C.GoString(operationID), C.GoString(conversationID), int64(msgDestructTime))
}

//export set_conversation_is_msg_destruct
func set_conversation_is_msg_destruct(operationID *C.char, conversationID *C.char, isMsgDestruct C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetConversationIsMsgDestruct(baseCallback, C.GoString(operationID), C.GoString(conversationID), parseBool(int(isMsgDestruct)))
}

//export hide_conversation
func hide_conversation(operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.HideConversation(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export get_conversation_recv_message_opt
func get_conversation_recv_message_opt(operationID *C.char, conversationIDList *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_GetConversationRecvMessageOptResp_List)
	open_im_sdk.GetConversationRecvMessageOpt(baseCallback, C.GoString(operationID), C.GoString(conversationIDList))
}

//export set_conversation_draft
func set_conversation_draft(operationID *C.char, conversationID *C.char, draftText *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetConversationDraft(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(draftText))
}

//export reset_conversation_group_at_type
func reset_conversation_group_at_type(operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.ResetConversationGroupAtType(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export pin_conversation
func pin_conversation(operationID *C.char, conversationID *C.char, isPinned C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.PinConversation(baseCallback, C.GoString(operationID), C.GoString(conversationID), parseBool(int(isPinned)))
}

//export set_conversation_private_chat
func set_conversation_private_chat(operationID *C.char, conversationID *C.char, isPrivate C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetConversationPrivateChat(baseCallback, C.GoString(operationID), C.GoString(conversationID),
		parseBool(int(isPrivate)))
}

//export set_conversation_burn_duration
func set_conversation_burn_duration(operationID *C.char, conversationID *C.char, duration C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetConversationBurnDuration(baseCallback, C.GoString(operationID), C.GoString(conversationID), int32(duration))
}

//export set_conversation_recv_message_opt
func set_conversation_recv_message_opt(operationID *C.char, conversationID *C.char, opt C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetConversationRecvMessageOpt(baseCallback, C.GoString(operationID), C.GoString(conversationID), int(opt))
}

//export get_total_unread_msg_count
func get_total_unread_msg_count(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Int)
	open_im_sdk.GetTotalUnreadMsgCount(baseCallback, C.GoString(operationID))
}

//export get_at_all_tag
func get_at_all_tag(operationID *C.char) *C.char {
	return C.CString(open_im_sdk.GetAtAllTag(C.GoString(operationID)))
}

//export get_conversation_id_by_session_type
func get_conversation_id_by_session_type(operationID *C.char, sourceID *C.char, sessionType C.int) *C.char {
	return C.CString(open_im_sdk.GetConversationIDBySessionType(C.GoString(operationID), C.GoString(sourceID), int(sessionType)))
}

//export send_message
func send_message(operationID, message, recvID, groupID, offlinePushInfo *C.char) {
	sendMsgCallback := NewSendMessageCallback(operationID)
	open_im_sdk.SendMessage(sendMsgCallback, C.GoString(operationID), C.GoString(message), C.GoString(recvID),
		C.GoString(groupID), C.GoString(offlinePushInfo))
}

//export send_message_not_oss
func send_message_not_oss(operationID, message, recvID, groupID, offlinePushInfo *C.char) {
	sendMsgCallback := NewSendMessageCallback(operationID)
	open_im_sdk.SendMessageNotOss(sendMsgCallback, C.GoString(operationID), C.GoString(message), C.GoString(recvID),
		C.GoString(groupID), C.GoString(offlinePushInfo))
}

//export find_message_list
func find_message_list(operationID *C.char, findMessageOptions *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_FindMessageListCallback)
	open_im_sdk.FindMessageList(baseCallback, C.GoString(operationID), C.GoString(findMessageOptions))
}

//export get_advanced_history_message_list
func get_advanced_history_message_list(operationID, getMessageOptions *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_GetAdvancedHistoryMessageListParams)
	open_im_sdk.GetAdvancedHistoryMessageList(baseCallback, C.GoString(operationID), C.GoString(getMessageOptions))
}

//export get_advanced_history_message_list_reverse
func get_advanced_history_message_list_reverse(operationID *C.char, getMessageOptions *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_GetAdvancedHistoryMessageListParams)
	open_im_sdk.GetAdvancedHistoryMessageListReverse(baseCallback, C.GoString(operationID), C.GoString(getMessageOptions))
}

//export revoke_message
func revoke_message(operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_GetAdvancedHistoryMessageListParams)
	open_im_sdk.RevokeMessage(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

//export typing_status_update
func typing_status_update(operationID *C.char, recvID *C.char, msgTip *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.TypingStatusUpdate(baseCallback, C.GoString(operationID), C.GoString(recvID), C.GoString(msgTip))
}

//export mark_conversation_message_as_read
func mark_conversation_message_as_read(operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.MarkConversationMessageAsRead(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export delete_message_from_local_storage
func delete_message_from_local_storage(operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.DeleteMessageFromLocalStorage(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

//export delete_message
func delete_message(operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.DeleteMessage(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

//export hide_all_conversations
func hide_all_conversations(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.HideAllConversations(baseCallback, C.GoString(operationID))
}

//export delete_all_msg_from_local_and_svr
func delete_all_msg_from_local_and_svr(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.DeleteAllMsgFromLocalAndSvr(baseCallback, C.GoString(operationID))
}

//export delete_all_msg_from_local
func delete_all_msg_from_local(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.DeleteAllMsgFromLocal(baseCallback, C.GoString(operationID))
}

//export clear_conversation_and_delete_all_msg
func clear_conversation_and_delete_all_msg(operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.ClearConversationAndDeleteAllMsg(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export delete_conversation_and_delete_all_msg
func delete_conversation_and_delete_all_msg(operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.DeleteConversationAndDeleteAllMsg(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export insert_single_message_to_local_storage
func insert_single_message_to_local_storage(operationID *C.char, message *C.char, recvID *C.char, sendID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_MsgStruct)
	open_im_sdk.InsertSingleMessageToLocalStorage(baseCallback, C.GoString(operationID), C.GoString(message), C.GoString(recvID), C.GoString(sendID))
}

//export insert_group_message_to_local_storage
func insert_group_message_to_local_storage(operationID *C.char, message *C.char, groupID *C.char, sendID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_MsgStruct)
	open_im_sdk.InsertGroupMessageToLocalStorage(baseCallback, C.GoString(operationID), C.GoString(message), C.GoString(groupID), C.GoString(sendID))
}

//export search_local_messages
func search_local_messages(operationID *C.char, searchParam *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_SearchLocalMessagesCallback)
	open_im_sdk.SearchLocalMessages(baseCallback, C.GoString(operationID), C.GoString(searchParam))
}

//export set_message_local_ex
func set_message_local_ex(operationID *C.char, conversationID *C.char, clientMsgID *C.char, localEx *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetMessageLocalEx(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID), C.GoString(localEx))
}

// =====================================================user===============================================

//export get_users_info
func get_users_info(operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_FullUserInfo_List)
	open_im_sdk.GetUsersInfo(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

//export get_users_info_with_cache
func get_users_info_with_cache(operationID *C.char, userIDs *C.char, groupID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_FullUserInfoWithCache_List)
	open_im_sdk.GetUsersInfoWithCache(baseCallback, C.GoString(operationID), C.GoString(userIDs), C.GoString(groupID))
}

//export get_users_info_from_srv
func get_users_info_from_srv(operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalUser_List)
	open_im_sdk.GetUsersInfoFromSrv(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

//export set_self_info
func set_self_info(operationID *C.char, userInfo *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetSelfInfo(baseCallback, C.GoString(operationID), C.GoString(userInfo))
}

//export set_global_recv_message_opt
func set_global_recv_message_opt(operationID *C.char, opt C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetGlobalRecvMessageOpt(baseCallback, C.GoString(operationID), int(opt))
}

//export get_self_user_info
func get_self_user_info(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalUser)
	open_im_sdk.GetSelfUserInfo(baseCallback, C.GoString(operationID))
}

//export update_msg_sender_info
func update_msg_sender_info(operationID *C.char, nickname *C.char, faceURL *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.UpdateMsgSenderInfo(baseCallback, C.GoString(operationID), C.GoString(nickname), C.GoString(faceURL))
}

//export subscribe_users_status
func subscribe_users_status(operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_OnlineStatus_List)
	open_im_sdk.SubscribeUsersStatus(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

//export unsubscribe_users_status
func unsubscribe_users_status(operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.UnsubscribeUsersStatus(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

//export get_subscribe_users_status
func get_subscribe_users_status(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_OnlineStatus_List)
	open_im_sdk.GetSubscribeUsersStatus(baseCallback, C.GoString(operationID))
}

//export get_user_status
func get_user_status(operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_OnlineStatus_List)
	open_im_sdk.GetUserStatus(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

// =====================================================friend===============================================
//
//export get_specified_friends_info
func get_specified_friends_info(operationID *C.char, userIDList *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_FullUserInfo_List)
	open_im_sdk.GetSpecifiedFriendsInfo(baseCallback, C.GoString(operationID), C.GoString(userIDList))
}

//export get_friend_list
func get_friend_list(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_FullUserInfo_List)
	open_im_sdk.GetFriendList(baseCallback, C.GoString(operationID))
}

//export get_friend_list_page
func get_friend_list_page(operationID *C.char, offset C.int, count C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_FullUserInfo_List)
	open_im_sdk.GetFriendListPage(baseCallback, C.GoString(operationID), int32(offset), int32(count))
}

//export search_friends
func search_friends(operationID *C.char, searchParam *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_SearchFriendItem_List)
	open_im_sdk.SearchFriends(baseCallback, C.GoString(operationID), C.GoString(searchParam))
}

//export check_friend
func check_friend(operationID *C.char, userIDList *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_UserIDResult_List)
	open_im_sdk.CheckFriend(baseCallback, C.GoString(operationID), C.GoString(userIDList))
}

//export add_friend
func add_friend(operationID *C.char, userIDReqMsg *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.AddFriend(baseCallback, C.GoString(operationID), C.GoString(userIDReqMsg))
}

//export set_friend_remark
func set_friend_remark(operationID *C.char, userIDRemark *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetFriendRemark(baseCallback, C.GoString(operationID), C.GoString(userIDRemark))
}

//export delete_friend
func delete_friend(operationID *C.char, friendUserID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.DeleteFriend(baseCallback, C.GoString(operationID), C.GoString(friendUserID))
}

//export get_friend_application_list_as_recipient
func get_friend_application_list_as_recipient(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalFriendRequest_List)
	open_im_sdk.GetFriendApplicationListAsRecipient(baseCallback, C.GoString(operationID))
}

//export get_friend_application_list_as_applicant
func get_friend_application_list_as_applicant(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalFriendRequest_List)
	open_im_sdk.GetFriendApplicationListAsApplicant(baseCallback, C.GoString(operationID))
}

//export accept_friend_application
func accept_friend_application(operationID *C.char, userIDHandleMsg *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.AcceptFriendApplication(baseCallback, C.GoString(operationID), C.GoString(userIDHandleMsg))
}

//export refuse_friend_application
func refuse_friend_application(operationID *C.char, userIDHandleMsg *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.RefuseFriendApplication(baseCallback, C.GoString(operationID), C.GoString(userIDHandleMsg))
}

//export add_black
func add_black(operationID *C.char, blackUserID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.AddBlack(baseCallback, C.GoString(operationID), C.GoString(blackUserID))
}

//export get_black_list
func get_black_list(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalBlack_List)
	open_im_sdk.GetBlackList(baseCallback, C.GoString(operationID))
}

//export remove_black
func remove_black(operationID *C.char, removeUserID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.RemoveBlack(baseCallback, C.GoString(operationID), C.GoString(removeUserID))
}

// =====================================================group===============================================
// CreateGroup creates a group
//
//export create_group
func create_group(operationID, cGroupReqInfo *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_GroupInfo)
	open_im_sdk.CreateGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupReqInfo))
}

// JoinGroup joins a group
//
//export join_group
func join_group(operationID, cGroupID, cReqMsg *C.char, cJoinSource C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.JoinGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupID), C.GoString(cReqMsg),
		int32(cJoinSource))
}

// QuitGroup quits a group
//
//export quit_group
func quit_group(operationID, cGroupID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.QuitGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupID))
}

// DismissGroup dismisses a group
//
//export dismiss_group
func dismiss_group(operationID, cGroupID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.DismissGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupID))
}

// ChangeGroupMute changes the mute status of a group
//
//export change_group_mute
func change_group_mute(operationID, cGroupID *C.char, cIsMute C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.ChangeGroupMute(baseCallback, C.GoString(operationID), C.GoString(cGroupID), parseBool(int(cIsMute)))
}

// ChangeGroupMemberMute changes the mute status of a group member
//
//export change_group_member_mute
func change_group_member_mute(operationID, cGroupID, cUserID *C.char, cMutedSeconds C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.ChangeGroupMemberMute(baseCallback, C.GoString(operationID), C.GoString(cGroupID), C.GoString(cUserID),
		int(cMutedSeconds))
}

// SetGroupMemberRoleLevel sets the role level of a group member
//
//export set_group_member_role_level
func set_group_member_role_level(operationID, cGroupID, cUserID *C.char, cRoleLevel C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetGroupMemberRoleLevel(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cUserID), int(cRoleLevel))
}

// SetGroupMemberInfo sets the information of a group member
//
//export set_group_member_info
func set_group_member_info(operationID *C.char, cGroupMemberInfo *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetGroupMemberInfo(baseCallback, C.GoString(operationID), C.GoString(cGroupMemberInfo))
}

// GetJoinedGroupList retrieves the list of joined groups
//
//export get_joined_group_list
func get_joined_group_list(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalGroup_List)
	open_im_sdk.GetJoinedGroupList(baseCallback, C.GoString(operationID))
}

// GetSpecifiedGroupsInfo retrieves the information of specified groups
//
//export get_specified_groups_info
func get_specified_groups_info(operationID, cGroupIDList *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalGroup_List)
	open_im_sdk.GetSpecifiedGroupsInfo(baseCallback, C.GoString(operationID), C.GoString(cGroupIDList))
}

// SearchGroups searches for groups
//
//export search_groups
func search_groups(operationID, cSearchParam *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalGroup_List)
	open_im_sdk.SearchGroups(baseCallback, C.GoString(operationID), C.GoString(cSearchParam))
}

// SetGroupInfo sets the information of a group
//
//export set_group_info
func set_group_info(operationID, cGroupInfo *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetGroupInfo(baseCallback, C.GoString(operationID), C.GoString(cGroupInfo))
}

// SetGroupVerification sets the verification mode of a group
//
//export set_group_verification
func set_group_verification(operationID, cGroupID *C.char, cVerification C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetGroupVerification(baseCallback, C.GoString(operationID), C.GoString(cGroupID), int32(cVerification))
}

// SetGroupLookMemberInfo sets the member information visibility of a group
//
//export set_group_look_member_info
func set_group_look_member_info(operationID, cGroupID *C.char, cRule C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetGroupLookMemberInfo(baseCallback, C.GoString(operationID), C.GoString(cGroupID), int32(cRule))
}

// SetGroupApplyMemberFriend sets the friend rule for group applicants
//
//export set_group_apply_member_friend
func set_group_apply_member_friend(operationID, cGroupID *C.char, cRule C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetGroupApplyMemberFriend(baseCallback, C.GoString(operationID), C.GoString(cGroupID), int32(cRule))
}

// GetGroupMemberList retrieves the list of group members
//
//export get_group_member_list
func get_group_member_list(operationID, cGroupID *C.char, cFilter, cOffset, cCount C.int) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalGroupMember_List)
	open_im_sdk.GetGroupMemberList(baseCallback, C.GoString(operationID), C.GoString(cGroupID), int32(cFilter),
		int32(cOffset), int32(cCount))
}

// GetGroupMemberOwnerAndAdmin retrieves the owner and admin members of a group
//
//export get_group_member_owner_and_admin
func get_group_member_owner_and_admin(operationID, cGroupID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalGroupMember_List)
	open_im_sdk.GetGroupMemberOwnerAndAdmin(baseCallback, C.GoString(operationID), C.GoString(cGroupID))
}

// GetGroupMemberListByJoinTimeFilter retrieves the list of group members filtered by join time
//
//export get_group_member_list_by_join_time_filter
func get_group_member_list_by_join_time_filter(operationID, cGroupID *C.char, cOffset,
	cCount C.int, cJoinTimeBegin, cJoinTimeEnd C.longlong, cFilterUserIDList *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalGroupMember_List)
	open_im_sdk.GetGroupMemberListByJoinTimeFilter(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		int32(cOffset), int32(cCount), int64(cJoinTimeBegin), int64(cJoinTimeEnd), C.GoString(cFilterUserIDList))
}

// GetSpecifiedGroupMembersInfo retrieves the information of specified group members
//
//export get_specified_group_members_info
func get_specified_group_members_info(operationID, cGroupID, cUserIDList *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalGroupMember_List)
	open_im_sdk.GetSpecifiedGroupMembersInfo(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cUserIDList))
}

// KickGroupMember kicks group members
//
//export kick_group_member
func kick_group_member(operationID, cGroupID, cReason, cUserIDList *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.KickGroupMember(baseCallback, C.GoString(operationID), C.GoString(cGroupID), C.GoString(cReason),
		C.GoString(cUserIDList))
}

// TransferGroupOwner transfers the ownership of a group
//
//export transfer_group_owner
func transfer_group_owner(operationID, cGroupID, cNewOwnerUserID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.TransferGroupOwner(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cNewOwnerUserID))
}

// InviteUserToGroup invites users to a group
//
//export invite_user_to_group
func invite_user_to_group(operationID, cGroupID, cReason, cUserIDList *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.InviteUserToGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupID), C.GoString(cReason),
		C.GoString(cUserIDList))
}

// GetGroupApplicationListAsRecipient retrieves the group application list as a recipient
//
//export get_group_application_list_as_recipient
func get_group_application_list_as_recipient(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalAdminGroupRequest_List)
	open_im_sdk.GetGroupApplicationListAsRecipient(baseCallback, C.GoString(operationID))
}

// GetGroupApplicationListAsApplicant retrieves the group application list as an applicant
//
//export get_group_application_list_as_applicant
func get_group_application_list_as_applicant(operationID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalGroupRequest_List)
	open_im_sdk.GetGroupApplicationListAsApplicant(baseCallback, C.GoString(operationID))
}

// AcceptGroupApplication accepts a group application
//
//export accept_group_application
func accept_group_application(operationID, cGroupID, cFromUserID, cHandleMsg *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.AcceptGroupApplication(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cFromUserID), C.GoString(cHandleMsg))
}

// RefuseGroupApplication refuses a group application
//
//export refuse_group_application
func refuse_group_application(operationID, cGroupID, cFromUserID, cHandleMsg *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.RefuseGroupApplication(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cFromUserID), C.GoString(cHandleMsg))
}

// SetGroupMemberNickname sets the nickname of a group member
//
//export set_group_member_nickname
func set_group_member_nickname(operationID, cGroupID, cUserID, cGroupMemberNickname *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Empty)
	open_im_sdk.SetGroupMemberNickname(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cUserID), C.GoString(cGroupMemberNickname))
}

// SearchGroupMembers searches for group members
//
//export search_group_members
func search_group_members(operationID, cSearchParam *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_LocalGroupMember_List)
	open_im_sdk.SearchGroupMembers(baseCallback, C.GoString(operationID), C.GoString(cSearchParam))
}

// IsJoinGroup checks if the user has joined a group
//
//export is_join_group
func is_join_group(operationID, cGroupID *C.char) {
	baseCallback := NewBaseCallback(operationID, DataType_Bool)
	open_im_sdk.IsJoinGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupID))
}

func main() {

}
