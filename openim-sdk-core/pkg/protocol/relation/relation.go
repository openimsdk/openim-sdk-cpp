// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package relation

import (
	"errors"
	"fmt"
	"github.com/openimsdk/protocol/constant"
)

func (x *GetPaginationFriendsReq) Check() error {
	if x.Pagination == nil {
		return errors.New("pagination is empty")
	}
	if x.Pagination.PageNumber < 1 {
		return errors.New("pageNumber is invalid")
	}
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	return nil
}

func (x *ApplyToAddFriendReq) Check() error {
	if x.ToUserID == "" {
		return errors.New("toUserID is empty")
	}
	if x.FromUserID == "" {
		return errors.New("fromUserID is empty")
	}
	return nil
}

func (x *ImportFriendReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.FriendUserIDs == nil {
		return errors.New("friendUserIDS is empty")
	}

	if len(x.FriendUserIDs) > constant.ParamMaxLength {
		return errors.New("too many FriendUserIDs, need to be less than 1000")
	}
	return nil
}

func (x *GetPaginationFriendsApplyToReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	if x.Pagination == nil {
		return errors.New("pagination is empty")
	}
	if x.Pagination.PageNumber < 1 {
		return errors.New("pageNumber is invalid")
	}
	return nil
}

func (x *GetDesignatedFriendsReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.FriendUserIDs == nil {
		return errors.New("friendUserIDS is empty")
	}

	// if len(x.FriendUserIDs) > 1000 {
	// 	return errors.New("too many FriendUserIDs, need to be less than 1000")
	// }
	return nil
}

func (x *AddBlackReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.BlackUserID == "" {
		return errors.New("blackUserID is empty")
	}
	return nil
}

func (x *RemoveBlackReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.BlackUserID == "" {
		return errors.New("blackUserID is empty")
	}
	return nil
}

func (x *GetPaginationBlacksReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	if x.Pagination == nil {
		return errors.New("pagination is empty")
	}
	if x.Pagination.PageNumber < 1 {
		return errors.New("pageNumber is invalid")
	}
	return nil
}

func (x *IsFriendReq) Check() error {
	if x.UserID1 == "" {
		return errors.New("userID1 is empty")
	}
	if x.UserID2 == "" {
		return errors.New("userID2 is empty")
	}
	return nil
}

func (x *IsBlackReq) Check() error {
	if x.UserID1 == "" {
		return errors.New("userID1 is empty")
	}
	if x.UserID2 == "" {
		return errors.New("userID2 is empty")
	}
	return nil
}

func (x *DeleteFriendReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("userID1 is empty")
	}
	if x.FriendUserID == "" {
		return errors.New("userID2 is empty")
	}
	return nil
}

func (x *RespondFriendApplyReq) Check() error {
	if x.ToUserID == "" {
		return errors.New("toUserID is empty")
	}
	if x.FromUserID == "" {
		return errors.New("fromUserID is empty")
	}
	return nil
}

func (x *SetFriendRemarkReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.FriendUserID == "" {
		return errors.New("fromUserID is empty")
	}
	return nil
}

func (x *GetPaginationFriendsApplyFromReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	if x.Pagination == nil {
		return errors.New("pagination is empty")
	}
	if x.Pagination.PageNumber < 1 {
		return errors.New("pageNumber is invalid")
	}
	return nil
}

func (x *GetFriendIDsReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	return nil
}
func (x *GetDesignatedFriendsApplyReq) Check() error {
	if x.FromUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.ToUserID == "" {
		return errors.New("toUserID is empty")
	}
	return nil
}
func (x *UpdateFriendsReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.FriendUserIDs == nil {
		return errors.New("friendUserIDs is empty")
	}

	if len(x.FriendUserIDs) > constant.ParamMaxLength {
		return errors.New("too many FriendUserIDs, need to be less than 1000")
	}
	return nil
}
func (x *GetSpecifiedFriendsInfoReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.UserIDList == nil {
		return errors.New("userIDList is empty")
	}

	if len(x.UserIDList) > constant.ParamMaxLength {
		return errors.New("too many UserIDs, need to be less than 1000")
	}
	return nil
}
func (x *GetFullFriendUserIDsReq) Check() error {
	if len(x.UserID) == 0 {
		return errors.New("userID is empty")
	}
	return nil
}

func (x *GetPaginationFriendsApplyToResp) Format() any {
	if x.Total > 20 {
		return fmt.Sprintf("len is %v", x.Total)
	}
	return x
}

func (x *GetPaginationFriendsApplyFromResp) Format() any {
	if x.Total > 20 {
		return fmt.Sprintf("len is %v", x.Total)
	}
	return x
}

func (x *GetPaginationFriendsResp) Format() any {
	if len(x.FriendsInfo) > 20 {
		return fmt.Sprintf("len is %v", len(x.FriendsInfo))
	}
	return x
}
