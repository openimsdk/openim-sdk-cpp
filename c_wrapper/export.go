package main

/*
#include <stdio.h>
typedef void (*CB)();
typedef void (*CB_I_S)(int,char *);
typedef void (*CB_S)(char *);
typedef void (*CB_I)(int);
__attribute__((weak))
void Call_CB(CB func)
{
    func();
}
__attribute__((weak))
void Call_CB_I_S(CB_I_S func,int errCode,char* errMsg)
{
    func(errCode,errMsg);
}
__attribute__((weak))
void Call_CB_S(CB_S func,char* data)
{
    func(data);
}
__attribute__((weak))
void Call_CB_I(CB_I func,int progress)
{
    func(progress);
}
*/
import "C"

import (
	"open_im_sdk/open_im_sdk"
)

type ConnCallback struct {
	onConnecting       C.CB
	onConnectSuccess   C.CB
	onConnectFailed    C.CB_I_S
	onKickedOffline    C.CB
	onUserTokenExpired C.CB
}

func NewConnCallback(onConnecting C.CB, onConnectSuccess C.CB,
	onKickedOffline C.CB, onUserTokenExpired C.CB, onConnectFailed C.CB_I_S) *ConnCallback {
	return &ConnCallback{onConnecting: onConnecting, onConnectSuccess: onConnectSuccess,
		onKickedOffline: onKickedOffline, onUserTokenExpired: onUserTokenExpired, onConnectFailed: onConnectFailed}
}

func (c ConnCallback) OnConnecting() {
	C.Call_CB(c.onConnecting)
}

func (c ConnCallback) OnConnectSuccess() {
	C.Call_CB(c.onConnectSuccess)
}

func (c ConnCallback) OnConnectFailed(errCode int32, errMsg string) {
	C.Call_CB_I_S(c.onConnectFailed, C.int(errCode), C.CString(errMsg))

}

func (c ConnCallback) OnKickedOffline() {
	C.Call_CB(c.onKickedOffline)
}

func (c ConnCallback) OnUserTokenExpired() {
	C.Call_CB(c.onUserTokenExpired)
}

type SendMessageCallback struct {
	*BaseCallback
	progressFunc C.CB_I
}

func NewSendMessageCallback(baseCallback *BaseCallback, progressFunc C.CB_I) *SendMessageCallback {
	return &SendMessageCallback{BaseCallback: baseCallback, progressFunc: progressFunc}
}

func (s SendMessageCallback) OnProgress(progress int) {
	C.Call_CB_I(s.progressFunc, C.int(progress))
}

type BaseCallback struct {
	successFunc C.CB_S
	failedFunc  C.CB_I_S
}

func NewBaseCallback(successFunc C.CB_S, failedFunc C.CB_I_S) *BaseCallback {
	return &BaseCallback{successFunc: successFunc, failedFunc: failedFunc}
}

func (b BaseCallback) OnError(errCode int32, errMsg string) {
	C.Call_CB_I_S(b.failedFunc, C.int(errCode), C.CString(errMsg))
}

func (b BaseCallback) OnSuccess(data string) {
	C.Call_CB_S(b.successFunc, C.CString(data))
}

//export  init_sdk
func init_sdk(onConnecting C.CB,
	onConnectSuccess C.CB,
	onKickedOffline C.CB,
	onUserTokenExpired C.CB,
	onConnectFailed C.CB_I_S,
	operationID *C.char, config *C.char) bool {
	callback := NewConnCallback(onConnecting, onConnectSuccess, onKickedOffline, onUserTokenExpired, onConnectFailed)
	return open_im_sdk.InitSDK(callback, C.GoString(operationID), C.GoString(config))
}

//export un_init_sdk
func un_init_sdk(operationID *C.char) {
	open_im_sdk.UnInitSDK(C.GoString(operationID))
}

//export  login
func login(successFunc C.CB_S, failedFunc C.CB_I_S, operationID, uid, token *C.char) {
	baseCallback := NewBaseCallback(successFunc, failedFunc)
	open_im_sdk.Login(baseCallback, C.GoString(operationID), C.GoString(uid), C.GoString(token))
}

//export  logout
func logout(successFunc C.CB_S, failedFunc C.CB_I_S, operationID *C.char) {
	baseCallback := NewBaseCallback(successFunc, failedFunc)
	open_im_sdk.Logout(baseCallback, C.GoString(operationID))
}

//export  network_status_changed
func network_status_changed(successFunc C.CB_S, failedFunc C.CB_I_S, operationID *C.char) {
	baseCallback := NewBaseCallback(successFunc, failedFunc)
	open_im_sdk.NetworkStatusChanged(baseCallback, C.GoString(operationID))
}

//export  get_login_status
func get_login_status(operationID *C.char) int {
	return open_im_sdk.GetLoginStatus(C.GoString(operationID))
}

//export  get_login_user
func get_login_user() string {
	return open_im_sdk.GetLoginUserID()
}

//export  create_text_message
func create_text_message(operationID, text *C.char) string {
	return open_im_sdk.CreateTextMessage(C.GoString(operationID), C.GoString(text))
}

//export create_advanced_text_message
func create_advanced_text_message(operationID, text, messageEntityList *C.char) string {
	return open_im_sdk.CreateAdvancedTextMessage(C.GoString(operationID), C.GoString(text), C.GoString(messageEntityList))
}

//export create_text_at_message
func create_text_at_message(operationID, text, atUserList, atUsersInfo, message *C.char) string {
	return open_im_sdk.CreateTextAtMessage(C.GoString(operationID), C.GoString(text), C.GoString(atUserList), C.GoString(atUsersInfo), C.GoString(message))
}

//export create_location_message
func create_location_message(operationID, description *C.char, longitude, latitude C.double) string {
	return open_im_sdk.CreateLocationMessage(C.GoString(operationID), C.GoString(description), float64(longitude), float64(latitude))
}

//export create_custom_message
func create_custom_message(operationID, data, extension, description *C.char) string {
	return open_im_sdk.CreateCustomMessage(C.GoString(operationID), C.GoString(data), C.GoString(extension), C.GoString(description))
}

//export create_quote_message
func create_quote_message(operationID, text, message *C.char) string {
	return open_im_sdk.CreateQuoteMessage(C.GoString(operationID), C.GoString(text), C.GoString(message))
}

//export create_advanced_quote_message
func create_advanced_quote_message(operationID, text, message, messageEntityList *C.char) string {
	return open_im_sdk.CreateAdvancedQuoteMessage(C.GoString(operationID), C.GoString(text), C.GoString(message), C.GoString(messageEntityList))
}

//export create_card_message
func create_card_message(operationID, cardInfo *C.char) string {
	return open_im_sdk.CreateCardMessage(C.GoString(operationID), C.GoString(cardInfo))
}

//export create_video_message_from_full_path
func create_video_message_from_full_path(operationID, videoFullPath, videoType *C.char, duration C.longlong, snapshotFullPath *C.char) string {
	return open_im_sdk.CreateVideoMessageFromFullPath(C.GoString(operationID), C.GoString(videoFullPath), C.GoString(videoType), int64(duration), C.GoString(snapshotFullPath))
}

//export create_image_message_from_full_path
func create_image_message_from_full_path(operationID, imageFullPath *C.char) string {
	return open_im_sdk.CreateImageMessageFromFullPath(C.GoString(operationID), C.GoString(imageFullPath))
}

//export create_sound_message_from_full_path
func create_sound_message_from_full_path(operationID, soundPath *C.char, duration C.longlong) string {
	return open_im_sdk.CreateSoundMessageFromFullPath(C.GoString(operationID), C.GoString(soundPath), int64(duration))
}

//export create_file_message_from_full_path
func create_file_message_from_full_path(operationID, fileFullPath, fileName *C.char) string {
	return open_im_sdk.CreateFileMessageFromFullPath(C.GoString(operationID), C.GoString(fileFullPath), C.GoString(fileName))
}

//export create_image_message
func create_image_message(operationID, imagePath *C.char) string {
	return open_im_sdk.CreateImageMessage(C.GoString(operationID), C.GoString(imagePath))
}

//export create_image_message_by_url
func create_image_message_by_url(operationID, sourcePicture, bigPicture, snapshotPicture *C.char) string {
	return open_im_sdk.CreateImageMessageByURL(C.GoString(operationID), C.GoString(sourcePicture), C.GoString(bigPicture), C.GoString(snapshotPicture))
}

//export create_sound_message_by_url
func create_sound_message_by_url(operationID, soundBaseInfo *C.char) string {
	return open_im_sdk.CreateSoundMessageByURL(C.GoString(operationID), C.GoString(soundBaseInfo))
}

//export create_sound_message
func create_sound_message(operationID, soundPath *C.char, duration C.longlong) string {
	return open_im_sdk.CreateSoundMessage(C.GoString(operationID), C.GoString(soundPath), int64(duration))
}

//export create_video_message_by_url
func create_video_message_by_url(operationID, videoBaseInfo *C.char) string {
	return open_im_sdk.CreateVideoMessageByURL(C.GoString(operationID), C.GoString(videoBaseInfo))
}

//export create_video_message
func create_video_message(operationID, videoPath *C.char, videoType *C.char, duration C.longlong, snapshotPath *C.char) string {
	return open_im_sdk.CreateVideoMessage(C.GoString(operationID), C.GoString(videoPath), C.GoString(videoType), int64(duration), C.GoString(snapshotPath))
}

//export create_file_message_by_url
func create_file_message_by_url(operationID, fileBaseInfo *C.char) string {
	return open_im_sdk.CreateFileMessageByURL(C.GoString(operationID), C.GoString(fileBaseInfo))
}

//export create_file_message
func create_file_message(operationID, filePath, fileName *C.char) string {
	return open_im_sdk.CreateFileMessage(C.GoString(operationID), C.GoString(filePath), C.GoString(fileName))
}

//export create_merger_message
func create_merger_message(operationID, messageList, title, summaryList *C.char) string {
	return open_im_sdk.CreateMergerMessage(C.GoString(operationID), C.GoString(messageList), C.GoString(title), C.GoString(summaryList))
}

//export create_face_message
func create_face_message(operationID *C.char, index C.int, data *C.char) string {
	return open_im_sdk.CreateFaceMessage(C.GoString(operationID), int(index), C.GoString(data))
}

//export create_forward_message
func create_forward_message(operationID, m *C.char) string {
	return open_im_sdk.CreateForwardMessage(C.GoString(operationID), C.GoString(m))
}

//export send_message
func send_message(successFunc C.CB_S, failedFunc C.CB_I_S, progressFunc C.CB_I, operationID, message, recvID, groupID, offlinePushInfo *C.char) {
	baseCallback := NewBaseCallback(successFunc, failedFunc)
	sendMsgCallback := NewSendMessageCallback(baseCallback, progressFunc)
	open_im_sdk.SendMessage(sendMsgCallback, C.GoString(operationID), C.GoString(message), C.GoString(recvID), C.GoString(groupID), C.GoString(offlinePushInfo))
}

func main() {

}
