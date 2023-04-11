package handlers

import (
	"financial_management/consts"
	"financial_management/service"
	"financial_management/setting"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func AdminGetUser(c *gin.Context) {
	var (
		uid = c.GetInt(consts.UserID)
	)
	if uid != setting.Config.Admin.ID {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	userList := service.GetAllUser()
	if userList == nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, userList)
}
