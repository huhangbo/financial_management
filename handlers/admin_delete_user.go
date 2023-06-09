package handlers

import (
	"financial_management/consts"
	"financial_management/service"
	"financial_management/setting"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func AdminDeleteUser(c *gin.Context) {
	var (
		userIDs []int
		uid     = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&userIDs); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	if uid != setting.Config.Admin.ID {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	if err := service.DeleteUser(userIDs); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
