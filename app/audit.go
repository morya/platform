// Copyright (c) 2017 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"github.com/mattermost/platform/model"
)

func GetAudits(userId string, limit int) (model.Audits, *model.AppError) {
	if result := <-Srv.Store.Audit().Get(userId, limit); result.Err != nil {
		return nil, result.Err
	} else {
		return result.Data.(model.Audits), nil
	}
}
