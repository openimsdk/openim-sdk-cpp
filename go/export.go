package main

/*
#include <stdio.h>
typedef void (*CB_S)(char *);
typedef void (*CB_I_S)(int,char *);
typedef void (*CB_S_I_S_S)(char *,int,char *,char *);
typedef void (*CB_S_I_S_S_I)(char *,int,char *,char *,int);
extern void Call_CB_S(CB_S func,char* data);
extern void Call_CB_I_S(CB_I_S func,int event,char* data);
extern void Call_CB_S_I_S_S(CB_S_I_S_S func,char *,int errCode,char* errMsg,char* data);
extern void Call_CB_S_I_S_S_I(CB_S_I_S_S_I func,char *,int errCode,char* errMsg,char* data,int progress);
extern CB_S DebugPrint;

*/
import "C"

import (
	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"
)

//export set_print
func set_print(print C.CB_S) {
	C.DebugPrint = print
}

func DebugPrint(info string) {
	C.Call_CB_S(C.DebugPrint, C.CString("DLL:"+info))
}

type Base struct {
	ErrCode int32  `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

// =====================================================listener===============================================

type ConnCallback struct {
	cCallback C.CB_I_S
}

func NewConnCallback(cCallback C.CB_I_S) *ConnCallback {
	return &ConnCallback{cCallback: cCallback}
}

func (c ConnCallback) OnConnecting() {
	C.Call_CB_I_S(c.cCallback, CONNECTING, NO_DATA)
}

func (c ConnCallback) OnConnectSuccess() {
	C.Call_CB_I_S(c.cCallback, CONNECT_SUCCESS, NO_DATA)
}

func (c ConnCallback) OnConnectFailed(errCode int32, errMsg string) {

	C.Call_CB_I_S(c.cCallback, CONNECT_FAILED, C.CString(StructToJsonString(Base{ErrCode: errCode, ErrMsg: errMsg})))

}

func (c ConnCallback) OnKickedOffline() {
	C.Call_CB_I_S(c.cCallback, KICKED_OFFLINE, NO_DATA)
}

func (c ConnCallback) OnUserTokenExpired() {
	C.Call_CB_I_S(c.cCallback, USER_TOKEN_EXPIRED, NO_DATA)
}

func (c ConnCallback) OnUserTokenInvalid(errMsg string) {
	C.Call_CB_I_S(c.cCallback, USER_TOKEN_INVALID, C.CString(errMsg))
}

type ConversationCallback struct {
	cCallback C.CB_I_S
}

func NewConversationCallback(cCallback C.CB_I_S) *ConversationCallback {
	return &ConversationCallback{cCallback: cCallback}
}

func (c ConversationCallback) OnSyncServerStart(reinstalled bool) {
	m := make(map[string]any)
	m["reinstalled"] = reinstalled
	C.Call_CB_I_S(c.cCallback, SYNC_SERVER_START, C.CString(StructToJsonString(m)))
}

func (c ConversationCallback) OnSyncServerProgress(progress int) {
	m := make(map[string]any)
	m["progress"] = progress
	C.Call_CB_I_S(c.cCallback, SYNC_SERVER_PROGRESS, C.CString(StructToJsonString(m)))
}

func (c ConversationCallback) OnSyncServerFinish(reinstalled bool) {
	m := make(map[string]any)
	m["reinstalled"] = reinstalled
	C.Call_CB_I_S(c.cCallback, SYNC_SERVER_FINISH, C.CString(StructToJsonString(m)))
}

func (c ConversationCallback) OnSyncServerFailed(reinstalled bool) {
	m := make(map[string]any)
	m["reinstalled"] = reinstalled
	C.Call_CB_I_S(c.cCallback, SYNC_SERVER_FAILED, C.CString(StructToJsonString(m)))
}

func (c ConversationCallback) OnNewConversation(conversationList string) {
	C.Call_CB_I_S(c.cCallback, NEW_CONVERSATION, C.CString(conversationList))
}

func (c ConversationCallback) OnConversationChanged(conversationList string) {
	C.Call_CB_I_S(c.cCallback, CONVERSATION_CHANGED, C.CString(conversationList))
}

func (c ConversationCallback) OnTotalUnreadMessageCountChanged(totalUnreadCount int32) {
	C.Call_CB_I_S(c.cCallback, TOTAL_UNREAD_MESSAGE_COUNT_CHANGED, C.CString(IntToString(totalUnreadCount)))
}

func (c ConversationCallback) OnConversationUserInputStatusChanged(change string) {
	C.Call_CB_I_S(c.cCallback, CONVERSATION_USER_INPUT_STATUS_CHANGED, C.CString(change))
}

type AdvancedMsgCallback struct {
	cCallback C.CB_I_S
}

func NewAdvancedMsgCallback(cCallback C.CB_I_S) *AdvancedMsgCallback {
	return &AdvancedMsgCallback{cCallback: cCallback}
}

func (a AdvancedMsgCallback) OnRecvNewMessage(message string) {
	C.Call_CB_I_S(a.cCallback, RECV_NEW_MESSAGE, C.CString(message))
}

func (a AdvancedMsgCallback) OnRecvC2CReadReceipt(msgReceiptList string) {
	C.Call_CB_I_S(a.cCallback, RECV_C2C_READ_RECEIPT, C.CString(msgReceiptList))
}

func (a AdvancedMsgCallback) OnRecvGroupReadReceipt(groupMsgReceiptList string) {
	C.Call_CB_I_S(a.cCallback, RECV_GROUP_READ_RECEIPT, C.CString(groupMsgReceiptList))
}

func (a AdvancedMsgCallback) OnNewRecvMessageRevoked(messageRevoked string) {
	C.Call_CB_I_S(a.cCallback, NEW_RECV_MESSAGE_REVOKED, C.CString(messageRevoked))
}

func (a AdvancedMsgCallback) OnRecvMessageExtensionsChanged(msgID string, reactionExtensionList string) {
	m := make(map[string]any)
	m["msgID"] = msgID
	m["reactionExtensionList"] = reactionExtensionList
	C.Call_CB_I_S(a.cCallback, RECV_MESSAGE_EXTENSIONS_CHANGED, C.CString(StructToJsonString(m)))
}

func (a AdvancedMsgCallback) OnRecvMessageExtensionsDeleted(msgID string, reactionExtensionKeyList string) {
	m := make(map[string]any)
	m["msgID"] = msgID
	m["reactionExtensionKeyList"] = reactionExtensionKeyList
	C.Call_CB_I_S(a.cCallback, RECV_MESSAGE_EXTENSIONS_DELETED, C.CString(StructToJsonString(m)))
}

func (a AdvancedMsgCallback) OnRecvMessageExtensionsAdded(msgID string, reactionExtensionList string) {
	m := make(map[string]any)
	m["msgID"] = msgID
	m["reactionExtensionList"] = reactionExtensionList
	C.Call_CB_I_S(a.cCallback, RECV_MESSAGE_EXTENSIONS_ADDED, C.CString(StructToJsonString(m)))
}

func (a AdvancedMsgCallback) OnRecvOfflineNewMessage(message string) {
	C.Call_CB_I_S(a.cCallback, RECV_OFFLINE_NEW_MESSAGE, C.CString(message))
}

func (a AdvancedMsgCallback) OnMsgDeleted(message string) {
	C.Call_CB_I_S(a.cCallback, MSG_DELETED, C.CString(message))
}

func (a AdvancedMsgCallback) OnRecvOnlineOnlyMessage(message string) {
	C.Call_CB_I_S(a.cCallback, RECV_ONLINE_ONLY_MESSAGE, C.CString(message))
}

type BatchMessageCallback struct {
	cCallback C.CB_I_S
}

func NewBatchMessageCallback(cCallback C.CB_I_S) *BatchMessageCallback {
	return &BatchMessageCallback{cCallback: cCallback}
}

func (b BatchMessageCallback) OnRecvNewMessages(messageList string) {
	C.Call_CB_I_S(b.cCallback, RECV_NEW_MESSAGES, C.CString(messageList))
}

func (b BatchMessageCallback) OnRecvOfflineNewMessages(messageList string) {
	C.Call_CB_I_S(b.cCallback, RECV_OFFLINE_NEW_MESSAGES, C.CString(messageList))
}

type FriendCallback struct {
	cCallback C.CB_I_S
}

func NewFriendCallback(cCallback C.CB_I_S) *FriendCallback {
	return &FriendCallback{cCallback: cCallback}
}

func (f FriendCallback) OnFriendApplicationAdded(friendApplication string) {
	C.Call_CB_I_S(f.cCallback, FRIEND_APPLICATION_ADDED, C.CString(friendApplication))
}

func (f FriendCallback) OnFriendApplicationDeleted(friendApplication string) {
	C.Call_CB_I_S(f.cCallback, FRIEND_APPLICATION_DELETED, C.CString(friendApplication))
}

func (f FriendCallback) OnFriendApplicationAccepted(friendApplication string) {
	C.Call_CB_I_S(f.cCallback, FRIEND_APPLICATION_ACCEPTED, C.CString(friendApplication))
}

func (f FriendCallback) OnFriendApplicationRejected(friendApplication string) {
	C.Call_CB_I_S(f.cCallback, FRIEND_APPLICATION_REJECTED, C.CString(friendApplication))
}

func (f FriendCallback) OnFriendAdded(friendInfo string) {
	C.Call_CB_I_S(f.cCallback, FRIEND_ADDED, C.CString(friendInfo))
}

func (f FriendCallback) OnFriendDeleted(friendInfo string) {
	C.Call_CB_I_S(f.cCallback, FRIEND_DELETED, C.CString(friendInfo))
}

func (f FriendCallback) OnFriendInfoChanged(friendInfo string) {
	C.Call_CB_I_S(f.cCallback, FRIEND_INFO_CHANGED, C.CString(friendInfo))
}

func (f FriendCallback) OnBlackAdded(blackInfo string) {
	C.Call_CB_I_S(f.cCallback, BLACK_ADDED, C.CString(blackInfo))
}

func (f FriendCallback) OnBlackDeleted(blackInfo string) {
	C.Call_CB_I_S(f.cCallback, BLACK_DELETED, C.CString(blackInfo))
}

type GroupCallback struct {
	cCallback C.CB_I_S
}

func NewGroupCallback(cCallback C.CB_I_S) *GroupCallback {
	return &GroupCallback{cCallback: cCallback}
}

func (g GroupCallback) OnJoinedGroupAdded(groupInfo string) {
	C.Call_CB_I_S(g.cCallback, JOINED_GROUP_ADDED, C.CString(groupInfo))
}

func (g GroupCallback) OnJoinedGroupDeleted(groupInfo string) {
	C.Call_CB_I_S(g.cCallback, JOINED_GROUP_DELETED, C.CString(groupInfo))
}

func (g GroupCallback) OnGroupMemberAdded(groupMemberInfo string) {
	C.Call_CB_I_S(g.cCallback, GROUP_MEMBER_ADDED, C.CString(groupMemberInfo))
}

func (g GroupCallback) OnGroupMemberDeleted(groupMemberInfo string) {
	C.Call_CB_I_S(g.cCallback, GROUP_MEMBER_DELETED, C.CString(groupMemberInfo))
}

func (g GroupCallback) OnGroupApplicationAdded(groupApplication string) {
	C.Call_CB_I_S(g.cCallback, GROUP_APPLICATION_ADDED, C.CString(groupApplication))
}

func (g GroupCallback) OnGroupApplicationDeleted(groupApplication string) {
	C.Call_CB_I_S(g.cCallback, GROUP_APPLICATION_DELETED, C.CString(groupApplication))
}

func (g GroupCallback) OnGroupInfoChanged(groupInfo string) {
	C.Call_CB_I_S(g.cCallback, GROUP_INFO_CHANGED, C.CString(groupInfo))
}

func (g GroupCallback) OnGroupDismissed(groupInfo string) {
	C.Call_CB_I_S(g.cCallback, GROUP_DISMISSED, C.CString(groupInfo))
}

func (g GroupCallback) OnGroupMemberInfoChanged(groupMemberInfo string) {
	C.Call_CB_I_S(g.cCallback, GROUP_MEMBER_INFO_CHANGED, C.CString(groupMemberInfo))
}

func (g GroupCallback) OnGroupApplicationAccepted(groupApplication string) {
	C.Call_CB_I_S(g.cCallback, GROUP_APPLICATION_ACCEPTED, C.CString(groupApplication))
}

func (g GroupCallback) OnGroupApplicationRejected(groupApplication string) {
	C.Call_CB_I_S(g.cCallback, GROUP_APPLICATION_REJECTED, C.CString(groupApplication))
}

type CustomBusinessCallback struct {
	cCallback C.CB_I_S
}

func NewCustomBusinessCallback(cCallback C.CB_I_S) *CustomBusinessCallback {
	return &CustomBusinessCallback{cCallback: cCallback}
}

func (c CustomBusinessCallback) OnRecvCustomBusinessMessage(businessMessage string) {
	C.Call_CB_I_S(c.cCallback, RECV_CUSTOM_BUSINESS_MESSAGE, C.CString(businessMessage))
}

type UserCallback struct {
	cCallback C.CB_I_S
}

func NewUserCallback(cCallback C.CB_I_S) *UserCallback {
	return &UserCallback{cCallback: cCallback}
}

func (u UserCallback) OnSelfInfoUpdated(userInfo string) {
	C.Call_CB_I_S(u.cCallback, SELF_INFO_UPDATED, C.CString(userInfo))
}

func (u UserCallback) OnUserStatusChanged(statusMap string) {
	C.Call_CB_I_S(u.cCallback, USER_STATUS_CHANGED, C.CString(statusMap))
}

func (u UserCallback) OnUserCommandAdd(userCommand string)    {}
func (u UserCallback) OnUserCommandDelete(userCommand string) {}
func (u UserCallback) OnUserCommandUpdate(userCommand string) {}

type SendMessageCallback struct {
	cCallback   C.CB_S_I_S_S_I
	operationID string
}

func NewSendMessageCallback(cCallback C.CB_S_I_S_S_I, operationID *C.char) *SendMessageCallback {
	return &SendMessageCallback{cCallback: cCallback, operationID: C.GoString(operationID)}
}

func (s SendMessageCallback) OnError(errCode int32, errMsg string) {
	C.Call_CB_S_I_S_S_I(s.cCallback, C.CString(s.operationID), C.int(errCode), C.CString(errMsg), NO_DATA, NO_PROGRESS)
}

func (s SendMessageCallback) OnSuccess(data string) {
	C.Call_CB_S_I_S_S_I(s.cCallback, C.CString(s.operationID), NO_ERR, NO_ERR_MSG, C.CString(data), NO_PROGRESS)
}

func (s SendMessageCallback) OnProgress(progress int) {
	C.Call_CB_S_I_S_S_I(s.cCallback, C.CString(s.operationID), NO_ERR, NO_ERR_MSG, NO_DATA, C.int(progress))
}

type BaseCallback struct {
	cCallback   C.CB_S_I_S_S
	operationID string
}

func NewBaseCallback(cCallback C.CB_S_I_S_S, operationID *C.char) *BaseCallback {
	return &BaseCallback{cCallback: cCallback, operationID: C.GoString(operationID)}
}

func (b BaseCallback) OnError(errCode int32, errMsg string) {
	C.Call_CB_S_I_S_S(b.cCallback, C.CString(b.operationID), C.int(errCode), C.CString(errMsg), NO_DATA)
}

func (b BaseCallback) OnSuccess(data string) {
	C.Call_CB_S_I_S_S(b.cCallback, C.CString(b.operationID), NO_ERR, NO_ERR_MSG, C.CString(data))
}

type UploadFileCallback struct {
	cCallback C.CB_I_S
}

func NewUploadFileCallback(cCallback C.CB_I_S) *UploadFileCallback {
	return &UploadFileCallback{cCallback: cCallback}
}

func (u UploadFileCallback) Open(size int64) {
	C.Call_CB_I_S(u.cCallback, OPEN, C.CString(IntToString(size)))
}

func (u UploadFileCallback) PartSize(partSize int64, num int) {
	m := make(map[string]any)
	m["partSize"] = partSize
	m["num"] = num
	C.Call_CB_I_S(u.cCallback, PART_SIZE, C.CString(StructToJsonString(m)))
}

func (u UploadFileCallback) HashPartProgress(index int, size int64, partHash string) {
	m := make(map[string]any)
	m["index"] = index
	m["size"] = size
	m["partHash"] = partHash
	C.Call_CB_I_S(u.cCallback, HASH_PART_PROGRESS, C.CString(StructToJsonString(m)))
}

func (u UploadFileCallback) HashPartComplete(partsHash string, fileHash string) {
	m := make(map[string]any)
	m["partsHash"] = partsHash
	m["fileHash"] = fileHash
	C.Call_CB_I_S(u.cCallback, HASH_PART_COMPLETE, C.CString(StructToJsonString(m)))
}

func (u UploadFileCallback) UploadID(uploadID string) {
	C.Call_CB_I_S(u.cCallback, UPLOAD_ID, C.CString(uploadID))
}

func (u UploadFileCallback) UploadPartComplete(index int, partSize int64, partHash string) {
	m := make(map[string]any)
	m["index"] = index
	m["partSize"] = partSize
	m["partHash"] = partHash
	C.Call_CB_I_S(u.cCallback, UPLOAD_PART_COMPLETE, C.CString(StructToJsonString(m)))
}

func (u UploadFileCallback) UploadComplete(fileSize int64, streamSize int64, storageSize int64) {
	m := make(map[string]any)
	m["fileSize"] = fileSize
	m["streamSize"] = streamSize
	m["storageSize"] = storageSize
	C.Call_CB_I_S(u.cCallback, UPLOAD_COMPLETE, C.CString(StructToJsonString(m)))
}

func (u UploadFileCallback) Complete(size int64, url string, typ int) {
	m := make(map[string]any)
	m["size"] = size
	m["url"] = url
	m["typ"] = typ
	C.Call_CB_I_S(u.cCallback, COMPLETE, C.CString(StructToJsonString(m)))
}

type UploadLogProgressCallback struct {
	cCallback C.CB_I_S
}

func NewUploadLogProgressCallback(cCallback C.CB_I_S) *UploadLogProgressCallback {
	return &UploadLogProgressCallback{cCallback: cCallback}
}

func (l UploadLogProgressCallback) OnProgress(current, size int64) {
	m := make(map[string]any)
	m["current"] = current
	m["size"] = size
	C.Call_CB_I_S(l.cCallback, ON_PROGRESS, C.CString(StructToJsonString(m)))
}

// =====================================================global_callback===============================================

//export set_group_listener
func set_group_listener(cCallback C.CB_I_S) {
	open_im_sdk.SetGroupListener(NewGroupCallback(cCallback))
}

//export set_conversation_listener
func set_conversation_listener(cCallback C.CB_I_S) {
	open_im_sdk.SetConversationListener(NewConversationCallback(cCallback))
}

//export set_advanced_msg_listener
func set_advanced_msg_listener(cCallback C.CB_I_S) {
	open_im_sdk.SetAdvancedMsgListener(NewAdvancedMsgCallback(cCallback))
}

//export set_batch_msg_listener
func set_batch_msg_listener(cCallback C.CB_I_S) {
	open_im_sdk.SetBatchMsgListener(NewBatchMessageCallback(cCallback))
}

//export set_user_listener
func set_user_listener(cCallback C.CB_I_S) {
	open_im_sdk.SetUserListener(NewUserCallback(cCallback))
}

//export set_friend_listener
func set_friend_listener(cCallback C.CB_I_S) {
	open_im_sdk.SetFriendListener(NewFriendCallback(cCallback))
}

//export set_custom_business_listener
func set_custom_business_listener(cCallback C.CB_I_S) {
	open_im_sdk.SetCustomBusinessListener(NewCustomBusinessCallback(cCallback))
}

////export set_messsage_kv_listener
//func set_messsage_kv_listener(cCallback C.CB_I_S) {
//	open_im_sdk.SetMessageKvInfoListener(NewMessageKVCallback(cCallback))
//}

// =====================================================init_login===============================================

//export  init_sdk
func init_sdk(
	cCallback C.CB_I_S,
	operationID *C.char, config *C.char) bool {
	callback := NewConnCallback(cCallback)
	return open_im_sdk.InitSDK(callback, C.GoString(operationID), C.GoString(config))
}

//export un_init_sdk
func un_init_sdk(operationID *C.char) {
	open_im_sdk.UnInitSDK(C.GoString(operationID))
}

//export  login
func login(cCallback C.CB_S_I_S_S, operationID, uid, token *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.Login(baseCallback, C.GoString(operationID), C.GoString(uid), C.GoString(token))
}

//export  logout
func logout(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.Logout(baseCallback, C.GoString(operationID))
}

//export set_app_background_status
func set_app_background_status(cCallback C.CB_S_I_S_S, operationID *C.char, isBackground C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetAppBackgroundStatus(baseCallback, C.GoString(operationID), parseBool(int(isBackground)))
}

//export  network_status_changed
func network_status_changed(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
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
func get_all_conversation_list(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetAllConversationList(baseCallback, C.GoString(operationID))
}

//export get_conversation_list_split
func get_conversation_list_split(cCallback C.CB_S_I_S_S, operationID *C.char, offset C.int, count C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetConversationListSplit(baseCallback, C.GoString(operationID), int(offset), int(count))
}

//export get_one_conversation
func get_one_conversation(cCallback C.CB_S_I_S_S, operationID *C.char, sessionType C.int, sourceID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetOneConversation(baseCallback, C.GoString(operationID), int32(sessionType), C.GoString(sourceID))
}

//export get_multiple_conversation
func get_multiple_conversation(cCallback C.CB_S_I_S_S, operationID *C.char, conversationIDList *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetMultipleConversation(baseCallback, C.GoString(operationID), C.GoString(conversationIDList))
}

//export set_conversation_msg_destruct_time
func set_conversation_msg_destruct_time(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, msgDestructTime C.longlong) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetConversationMsgDestructTime(baseCallback, C.GoString(operationID), C.GoString(conversationID), int64(msgDestructTime))
}

//export set_conversation_is_msg_destruct
func set_conversation_is_msg_destruct(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, isMsgDestruct C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetConversationIsMsgDestruct(baseCallback, C.GoString(operationID), C.GoString(conversationID), parseBool(int(isMsgDestruct)))
}

//export set_conversation_ex
func set_conversation_ex(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, ex *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetConversationEx(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(ex))
}

//export hide_conversation
func hide_conversation(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.HideConversation(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export get_conversation_recv_message_opt
func get_conversation_recv_message_opt(cCallback C.CB_S_I_S_S, operationID *C.char, conversationIDList *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetConversationRecvMessageOpt(baseCallback, C.GoString(operationID), C.GoString(conversationIDList))
}

//export set_conversation_draft
func set_conversation_draft(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, draftText *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetConversationDraft(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(draftText))
}

//export reset_conversation_group_at_type
func reset_conversation_group_at_type(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.ResetConversationGroupAtType(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export pin_conversation
func pin_conversation(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, isPinned C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.PinConversation(baseCallback, C.GoString(operationID), C.GoString(conversationID), parseBool(int(isPinned)))
}

//export set_conversation_private_chat
func set_conversation_private_chat(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, isPrivate C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetConversationPrivateChat(baseCallback, C.GoString(operationID), C.GoString(conversationID),
		parseBool(int(isPrivate)))
}

//export set_conversation_burn_duration
func set_conversation_burn_duration(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, duration C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetConversationBurnDuration(baseCallback, C.GoString(operationID), C.GoString(conversationID), int32(duration))
}

//export set_conversation_recv_message_opt
func set_conversation_recv_message_opt(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, opt C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetConversationRecvMessageOpt(baseCallback, C.GoString(operationID), C.GoString(conversationID), int(opt))
}

//export get_total_unread_msg_count
func get_total_unread_msg_count(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
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
func send_message(cCallback C.CB_S_I_S_S_I, operationID, message, recvID, groupID, offlinePushInfo *C.char, isOnlineOnly C.int) {
	sendMsgCallback := NewSendMessageCallback(cCallback, operationID)
	open_im_sdk.SendMessage(sendMsgCallback, C.GoString(operationID), C.GoString(message), C.GoString(recvID),
		C.GoString(groupID), C.GoString(offlinePushInfo), parseBool(int(isOnlineOnly)))
}

//export send_message_not_oss
func send_message_not_oss(cCallback C.CB_S_I_S_S_I, operationID, message, recvID, groupID, offlinePushInfo *C.char, isOnlineOnly C.int) {
	sendMsgCallback := NewSendMessageCallback(cCallback, operationID)
	open_im_sdk.SendMessageNotOss(sendMsgCallback, C.GoString(operationID), C.GoString(message), C.GoString(recvID),
		C.GoString(groupID), C.GoString(offlinePushInfo), parseBool(int(isOnlineOnly)))
}

//export find_message_list
func find_message_list(cCallback C.CB_S_I_S_S, operationID *C.char, findMessageOptions *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.FindMessageList(baseCallback, C.GoString(operationID), C.GoString(findMessageOptions))
}

//export get_advanced_history_message_list
func get_advanced_history_message_list(cCallback C.CB_S_I_S_S, operationID, getMessageOptions *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetAdvancedHistoryMessageList(baseCallback, C.GoString(operationID), C.GoString(getMessageOptions))
}

//export get_advanced_history_message_list_reverse
func get_advanced_history_message_list_reverse(cCallback C.CB_S_I_S_S, operationID *C.char, getMessageOptions *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetAdvancedHistoryMessageListReverse(baseCallback, C.GoString(operationID), C.GoString(getMessageOptions))
}

//export revoke_message
func revoke_message(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.RevokeMessage(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

//export typing_status_update
func typing_status_update(cCallback C.CB_S_I_S_S, operationID *C.char, recvID *C.char, msgTip *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.TypingStatusUpdate(baseCallback, C.GoString(operationID), C.GoString(recvID), C.GoString(msgTip))
}

//export mark_conversation_message_as_read
func mark_conversation_message_as_read(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.MarkConversationMessageAsRead(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export delete_message_from_local_storage
func delete_message_from_local_storage(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.DeleteMessageFromLocalStorage(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

//export delete_message
func delete_message(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, clientMsgID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.DeleteMessage(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

//export hide_all_conversations
func hide_all_conversations(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.HideAllConversations(baseCallback, C.GoString(operationID))
}

//export delete_all_msg_from_local_and_svr
func delete_all_msg_from_local_and_svr(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.DeleteAllMsgFromLocalAndSvr(baseCallback, C.GoString(operationID))
}

//export delete_all_msg_from_local
func delete_all_msg_from_local(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.DeleteAllMsgFromLocal(baseCallback, C.GoString(operationID))
}

//export clear_conversation_and_delete_all_msg
func clear_conversation_and_delete_all_msg(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.ClearConversationAndDeleteAllMsg(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export delete_conversation_and_delete_all_msg
func delete_conversation_and_delete_all_msg(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.DeleteConversationAndDeleteAllMsg(baseCallback, C.GoString(operationID), C.GoString(conversationID))
}

//export insert_single_message_to_local_storage
func insert_single_message_to_local_storage(cCallback C.CB_S_I_S_S, operationID *C.char, message *C.char, recvID *C.char, sendID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.InsertSingleMessageToLocalStorage(baseCallback, C.GoString(operationID), C.GoString(message), C.GoString(recvID), C.GoString(sendID))
}

//export insert_group_message_to_local_storage
func insert_group_message_to_local_storage(cCallback C.CB_S_I_S_S, operationID *C.char, message *C.char, groupID *C.char, sendID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.InsertGroupMessageToLocalStorage(baseCallback, C.GoString(operationID), C.GoString(message), C.GoString(groupID), C.GoString(sendID))
}

//export search_local_messages
func search_local_messages(cCallback C.CB_S_I_S_S, operationID *C.char, searchParam *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SearchLocalMessages(baseCallback, C.GoString(operationID), C.GoString(searchParam))
}

//export set_message_local_ex
func set_message_local_ex(cCallback C.CB_S_I_S_S, operationID *C.char, conversationID *C.char, clientMsgID *C.char, localEx *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetMessageLocalEx(baseCallback, C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID), C.GoString(localEx))
}

// =====================================================user===============================================

//export get_users_info
func get_users_info(cCallback C.CB_S_I_S_S, operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetUsersInfo(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

//export get_users_info_with_cache
func get_users_info_with_cache(cCallback C.CB_S_I_S_S, operationID *C.char, userIDs *C.char, groupID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetUsersInfoWithCache(baseCallback, C.GoString(operationID), C.GoString(userIDs), C.GoString(groupID))
}

//export get_users_info_from_srv
func get_users_info_from_srv(cCallback C.CB_S_I_S_S, operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetUsersInfoFromSrv(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

//export set_self_info
func set_self_info(cCallback C.CB_S_I_S_S, operationID *C.char, userInfo *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetSelfInfo(baseCallback, C.GoString(operationID), C.GoString(userInfo))
}

//export set_global_recv_message_opt
func set_global_recv_message_opt(cCallback C.CB_S_I_S_S, operationID *C.char, opt C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetGlobalRecvMessageOpt(baseCallback, C.GoString(operationID), int(opt))
}

//export get_self_user_info
func get_self_user_info(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetSelfUserInfo(baseCallback, C.GoString(operationID))
}

//export update_msg_sender_info
func update_msg_sender_info(cCallback C.CB_S_I_S_S, operationID *C.char, nickname *C.char, faceURL *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.UpdateMsgSenderInfo(baseCallback, C.GoString(operationID), C.GoString(nickname), C.GoString(faceURL))
}

// =====================================================file===============================================
//

//export upload_file
func upload_file(cCallback C.CB_S_I_S_S, operationID *C.char, req *C.char, uploadCallback C.CB_I_S) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	uploadFileCallback := NewUploadFileCallback(uploadCallback)
	open_im_sdk.UploadFile(baseCallback, C.GoString(operationID), C.GoString(req), uploadFileCallback)
}

// =====================================================friend===============================================
//
//export get_specified_friends_info
func get_specified_friends_info(cCallback C.CB_S_I_S_S, operationID *C.char, userIDList *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetSpecifiedFriendsInfo(baseCallback, C.GoString(operationID), C.GoString(userIDList))
}

//export get_friend_list
func get_friend_list(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetFriendList(baseCallback, C.GoString(operationID))
}

//export get_friend_list_page
func get_friend_list_page(cCallback C.CB_S_I_S_S, operationID *C.char, offset C.int, count C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetFriendListPage(baseCallback, C.GoString(operationID), int32(offset), int32(count))
}

//export search_friends
func search_friends(cCallback C.CB_S_I_S_S, operationID *C.char, searchParam *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SearchFriends(baseCallback, C.GoString(operationID), C.GoString(searchParam))
}

//export check_friend
func check_friend(cCallback C.CB_S_I_S_S, operationID *C.char, userIDList *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.CheckFriend(baseCallback, C.GoString(operationID), C.GoString(userIDList))
}

//export add_friend
func add_friend(cCallback C.CB_S_I_S_S, operationID *C.char, userIDReqMsg *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.AddFriend(baseCallback, C.GoString(operationID), C.GoString(userIDReqMsg))
}

//export set_friend_remark
func set_friend_remark(cCallback C.CB_S_I_S_S, operationID *C.char, userIDRemark *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetFriendRemark(baseCallback, C.GoString(operationID), C.GoString(userIDRemark))
}

//export delete_friend
func delete_friend(cCallback C.CB_S_I_S_S, operationID *C.char, friendUserID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.DeleteFriend(baseCallback, C.GoString(operationID), C.GoString(friendUserID))
}

//export get_friend_application_list_as_recipient
func get_friend_application_list_as_recipient(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetFriendApplicationListAsRecipient(baseCallback, C.GoString(operationID))
}

//export get_friend_application_list_as_applicant
func get_friend_application_list_as_applicant(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetFriendApplicationListAsApplicant(baseCallback, C.GoString(operationID))
}

//export accept_friend_application
func accept_friend_application(cCallback C.CB_S_I_S_S, operationID *C.char, userIDHandleMsg *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.AcceptFriendApplication(baseCallback, C.GoString(operationID), C.GoString(userIDHandleMsg))
}

//export refuse_friend_application
func refuse_friend_application(cCallback C.CB_S_I_S_S, operationID *C.char, userIDHandleMsg *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.RefuseFriendApplication(baseCallback, C.GoString(operationID), C.GoString(userIDHandleMsg))
}

//export add_black
func add_black(cCallback C.CB_S_I_S_S, operationID *C.char, blackUserID *C.char, ex *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.AddBlack(baseCallback, C.GoString(operationID), C.GoString(blackUserID), C.GoString(ex))
}

//export get_black_list
func get_black_list(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetBlackList(baseCallback, C.GoString(operationID))
}

//export remove_black
func remove_black(cCallback C.CB_S_I_S_S, operationID *C.char, removeUserID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.RemoveBlack(baseCallback, C.GoString(operationID), C.GoString(removeUserID))
}

//export set_friends_ex
func set_friends_ex(cCallback C.CB_S_I_S_S, operationID, friendIDs, ex *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetFriendsEx(baseCallback, C.GoString(operationID), C.GoString(friendIDs), C.GoString(ex))
}

// =====================================================group===============================================
// CreateGroup creates a group
//
//export create_group
func create_group(cCallback C.CB_S_I_S_S, operationID, cGroupReqInfo *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.CreateGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupReqInfo))
}

// JoinGroup joins a group
//
//export join_group
func join_group(cCallback C.CB_S_I_S_S, operationID, cGroupID, cReqMsg *C.char, cJoinSource C.int, ex *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.JoinGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupID), C.GoString(cReqMsg),
		int32(cJoinSource), C.GoString(ex))
}

// QuitGroup quits a group
//
//export quit_group
func quit_group(cCallback C.CB_S_I_S_S, operationID, cGroupID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.QuitGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupID))
}

// DismissGroup dismisses a group
//
//export dismiss_group
func dismiss_group(cCallback C.CB_S_I_S_S, operationID, cGroupID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.DismissGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupID))
}

// ChangeGroupMute changes the mute status of a group
//
//export change_group_mute
func change_group_mute(cCallback C.CB_S_I_S_S, operationID, cGroupID *C.char, cIsMute C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.ChangeGroupMute(baseCallback, C.GoString(operationID), C.GoString(cGroupID), parseBool(int(cIsMute)))
}

// ChangeGroupMemberMute changes the mute status of a group member
//
//export change_group_member_mute
func change_group_member_mute(cCallback C.CB_S_I_S_S, operationID, cGroupID, cUserID *C.char, cMutedSeconds C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.ChangeGroupMemberMute(baseCallback, C.GoString(operationID), C.GoString(cGroupID), C.GoString(cUserID),
		int(cMutedSeconds))
}

// SetGroupMemberRoleLevel sets the role level of a group member
//
//export set_group_member_role_level
func set_group_member_role_level(cCallback C.CB_S_I_S_S, operationID, cGroupID, cUserID *C.char, cRoleLevel C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetGroupMemberRoleLevel(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cUserID), int(cRoleLevel))
}

// SetGroupMemberInfo sets the information of a group member
//
//export set_group_member_info
func set_group_member_info(cCallback C.CB_S_I_S_S, operationID *C.char, cGroupMemberInfo *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetGroupMemberInfo(baseCallback, C.GoString(operationID), C.GoString(cGroupMemberInfo))
}

// GetJoinedGroupList retrieves the list of joined groups
//
//export get_joined_group_list
func get_joined_group_list(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetJoinedGroupList(baseCallback, C.GoString(operationID))
}

// GetJoinedGroupListPage retrieves the list of joined groups with pagination
//
//export get_joined_group_list_page
func get_joined_group_list_page(cCallback C.CB_S_I_S_S, operationID *C.char, offset, count C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetJoinedGroupListPage(baseCallback, C.GoString(operationID), int32(offset), int32(count))
}

// GetSpecifiedGroupsInfo retrieves the information of specified groups
//
//export get_specified_groups_info
func get_specified_groups_info(cCallback C.CB_S_I_S_S, operationID, cGroupIDList *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetSpecifiedGroupsInfo(baseCallback, C.GoString(operationID), C.GoString(cGroupIDList))
}

// SearchGroups searches for groups
//
//export search_groups
func search_groups(cCallback C.CB_S_I_S_S, operationID, cSearchParam *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SearchGroups(baseCallback, C.GoString(operationID), C.GoString(cSearchParam))
}

// SetGroupInfo sets the information of a group
//
//export set_group_info
func set_group_info(cCallback C.CB_S_I_S_S, operationID, cGroupInfo *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetGroupInfo(baseCallback, C.GoString(operationID), C.GoString(cGroupInfo))
}

// SetGroupVerification sets the verification mode of a group
//
//export set_group_verification
func set_group_verification(cCallback C.CB_S_I_S_S, operationID, cGroupID *C.char, cVerification C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetGroupVerification(baseCallback, C.GoString(operationID), C.GoString(cGroupID), int32(cVerification))
}

// SetGroupLookMemberInfo sets the member information visibility of a group
//
//export set_group_look_member_info
func set_group_look_member_info(cCallback C.CB_S_I_S_S, operationID, cGroupID *C.char, cRule C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetGroupLookMemberInfo(baseCallback, C.GoString(operationID), C.GoString(cGroupID), int32(cRule))
}

// SetGroupApplyMemberFriend sets the friend rule for group applicants
//
//export set_group_apply_member_friend
func set_group_apply_member_friend(cCallback C.CB_S_I_S_S, operationID, cGroupID *C.char, cRule C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetGroupApplyMemberFriend(baseCallback, C.GoString(operationID), C.GoString(cGroupID), int32(cRule))
}

// GetGroupMemberList retrieves the list of group members
//
//export get_group_member_list
func get_group_member_list(cCallback C.CB_S_I_S_S, operationID, cGroupID *C.char, cFilter, cOffset, cCount C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetGroupMemberList(baseCallback, C.GoString(operationID), C.GoString(cGroupID), int32(cFilter),
		int32(cOffset), int32(cCount))
}

// GetGroupMemberOwnerAndAdmin retrieves the owner and admin members of a group
//
//export get_group_member_owner_and_admin
func get_group_member_owner_and_admin(cCallback C.CB_S_I_S_S, operationID, cGroupID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetGroupMemberOwnerAndAdmin(baseCallback, C.GoString(operationID), C.GoString(cGroupID))
}

// GetGroupMemberListByJoinTimeFilter retrieves the list of group members filtered by join time
//
//export get_group_member_list_by_join_time_filter
func get_group_member_list_by_join_time_filter(cCallback C.CB_S_I_S_S, operationID, cGroupID *C.char, cOffset,
	cCount C.int, cJoinTimeBegin, cJoinTimeEnd C.longlong, cFilterUserIDList *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetGroupMemberListByJoinTimeFilter(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		int32(cOffset), int32(cCount), int64(cJoinTimeBegin), int64(cJoinTimeEnd), C.GoString(cFilterUserIDList))
}

// GetSpecifiedGroupMembersInfo retrieves the information of specified group members
//
//export get_specified_group_members_info
func get_specified_group_members_info(cCallback C.CB_S_I_S_S, operationID, cGroupID, cUserIDList *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetSpecifiedGroupMembersInfo(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cUserIDList))
}

// KickGroupMember kicks group members
//
//export kick_group_member
func kick_group_member(cCallback C.CB_S_I_S_S, operationID, cGroupID, cReason, cUserIDList *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.KickGroupMember(baseCallback, C.GoString(operationID), C.GoString(cGroupID), C.GoString(cReason),
		C.GoString(cUserIDList))
}

// TransferGroupOwner transfers the ownership of a group
//
//export transfer_group_owner
func transfer_group_owner(cCallback C.CB_S_I_S_S, operationID, cGroupID, cNewOwnerUserID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.TransferGroupOwner(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cNewOwnerUserID))
}

// InviteUserToGroup invites users to a group
//
//export invite_user_to_group
func invite_user_to_group(cCallback C.CB_S_I_S_S, operationID, cGroupID, cReason, cUserIDList *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.InviteUserToGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupID), C.GoString(cReason),
		C.GoString(cUserIDList))
}

// GetGroupApplicationListAsRecipient retrieves the group application list as a recipient
//
//export get_group_application_list_as_recipient
func get_group_application_list_as_recipient(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetGroupApplicationListAsRecipient(baseCallback, C.GoString(operationID))
}

// GetGroupApplicationListAsApplicant retrieves the group application list as an applicant
//
//export get_group_application_list_as_applicant
func get_group_application_list_as_applicant(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetGroupApplicationListAsApplicant(baseCallback, C.GoString(operationID))
}

// AcceptGroupApplication accepts a group application
//
//export accept_group_application
func accept_group_application(cCallback C.CB_S_I_S_S, operationID, cGroupID, cFromUserID, cHandleMsg *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.AcceptGroupApplication(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cFromUserID), C.GoString(cHandleMsg))
}

// RefuseGroupApplication refuses a group application
//
//export refuse_group_application
func refuse_group_application(cCallback C.CB_S_I_S_S, operationID, cGroupID, cFromUserID, cHandleMsg *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.RefuseGroupApplication(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cFromUserID), C.GoString(cHandleMsg))
}

// SetGroupMemberNickname sets the nickname of a group member
//
//export set_group_member_nickname
func set_group_member_nickname(cCallback C.CB_S_I_S_S, operationID, cGroupID, cUserID, cGroupMemberNickname *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetGroupMemberNickname(baseCallback, C.GoString(operationID), C.GoString(cGroupID),
		C.GoString(cUserID), C.GoString(cGroupMemberNickname))
}

// SearchGroupMembers searches for group members
//
//export search_group_members
func search_group_members(cCallback C.CB_S_I_S_S, operationID, cSearchParam *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SearchGroupMembers(baseCallback, C.GoString(operationID), C.GoString(cSearchParam))
}

// IsJoinGroup checks if the user has joined a group
//
//export is_join_group
func is_join_group(cCallback C.CB_S_I_S_S, operationID, cGroupID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.IsJoinGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupID))
}

// GetUsersInGroup retrieves the users in a group
//
//export get_users_in_group
func get_users_in_group(cCallback C.CB_S_I_S_S, operationID, cGroupID, userIDList *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetUsersInGroup(baseCallback, C.GoString(operationID), C.GoString(cGroupID), C.GoString(userIDList))
}

// =====================================================online===============================================

//export subscribe_users_status
func subscribe_users_status(cCallback C.CB_S_I_S_S, operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SubscribeUsersStatus(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

//export unsubscribe_users_status
func unsubscribe_users_status(cCallback C.CB_S_I_S_S, operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.UnsubscribeUsersStatus(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

//export get_subscribe_users_status
func get_subscribe_users_status(cCallback C.CB_S_I_S_S, operationID *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetSubscribeUsersStatus(baseCallback, C.GoString(operationID))
}

//export get_user_status
func get_user_status(cCallback C.CB_S_I_S_S, operationID *C.char, userIDs *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.GetUserStatus(baseCallback, C.GoString(operationID), C.GoString(userIDs))
}

// =====================================================third===============================================

//export update_fcm_token
func update_fcm_token(cCallback C.CB_S_I_S_S, operationID, fcmToken *C.char, expireTime C.longlong) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.UpdateFcmToken(baseCallback, C.GoString(operationID), C.GoString(fcmToken), int64(expireTime))
}

//export set_app_Badge
func set_app_Badge(cCallback C.CB_S_I_S_S, operationID *C.char, appUnreadCount C.int) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.SetAppBadge(baseCallback, C.GoString(operationID), int32(appUnreadCount))
}

//export upload_logs
func upload_logs(cCallback C.CB_S_I_S_S, operationID *C.char, line C.int, ex *C.char, uploadLogProgressCallback C.CB_I_S) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	uploadLogCallback := NewUploadLogProgressCallback(uploadLogProgressCallback)
	open_im_sdk.UploadLogs(baseCallback, C.GoString(operationID), int(line), C.GoString(ex), uploadLogCallback)
}

// export logs
func logs(cCallback C.CB_S_I_S_S, operationID *C.char, logLevel C.int, file *C.char, line C.int, msgs *C.char, err *C.char, keyAndValue *C.char) {
	baseCallback := NewBaseCallback(cCallback, operationID)
	open_im_sdk.Logs(baseCallback, C.GoString(operationID), int(logLevel), C.GoString(file), int(line), C.GoString(msgs), C.GoString(err), C.GoString(keyAndValue))
}

func main() {

}
