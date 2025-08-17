// Copyright © 2023 OpenIM SDK. All rights reserved.
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

//go:build !js
// +build !js

package db

import (
	"context"

	"github.com/openimsdk/openim-sdk-core/v3/pkg/db/model_struct"
	"github.com/openimsdk/tools/errs"
	"github.com/pkg/errors"
)

// ProcessUserCommandAdd adds a new user command to the database.
func (d *DataBase) ProcessUserCommandAdd(ctx context.Context, command *model_struct.LocalUserCommand) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()

	userCommand := model_struct.LocalUserCommand{
		UserID:     command.UserID,
		CreateTime: command.CreateTime,
		Type:       command.Type,
		Uuid:       command.Uuid,
		Value:      command.Value,
		Ex:         command.Ex,
	}

	return errs.WrapMsg(d.conn.WithContext(ctx).Create(&userCommand).Error, "ProcessUserCommandAdd failed")
}

// ProcessUserCommandUpdate updates an existing user command in the database.
func (d *DataBase) ProcessUserCommandUpdate(ctx context.Context, command *model_struct.LocalUserCommand) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()

	t := d.conn.WithContext(ctx).Model(command).Select("*").Updates(*command)
	if t.RowsAffected == 0 {
		return errs.WrapMsg(errors.New("RowsAffected == 0"), "no update")
	}
	return errs.WrapMsg(t.Error, "")

}

// ProcessUserCommandDelete deletes a user command from the database.
func (d *DataBase) ProcessUserCommandDelete(ctx context.Context, command *model_struct.LocalUserCommand) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	return errs.WrapMsg(d.conn.WithContext(ctx).Where("type = ? AND uuid = ?", command.Type, command.Uuid).Delete(&model_struct.LocalUserCommand{}).Error,
		"ProcessUserCommandDelete failed")
}

// ProcessUserCommandGetAll retrieves user commands from the database.
func (d *DataBase) ProcessUserCommandGetAll(ctx context.Context) ([]*model_struct.LocalUserCommand, error) {
	d.mRWMutex.RLock()
	defer d.mRWMutex.RUnlock()
	var commands []*model_struct.LocalUserCommand
	return commands, errs.WrapMsg(d.conn.WithContext(ctx).Find(&commands).Error, "ProcessUserCommandGetAll failed")
}
