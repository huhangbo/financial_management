package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/setting"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func UserUpdateInfo(c *gin.Context) {
	var (
		reqUser *model.User
		uid     = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&reqUser); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpUser := service.GetUserByID(reqUser.UserID)
	if tmpUser == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	if uid != setting.Config.Admin.ID && uid != reqUser.UserID {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	tmpUser.Username = reqUser.Username
	tmpUser.Gender = reqUser.Gender
	tmpUser.Email = reqUser.Email
	if err := service.UpdateUser(tmpUser); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, tmpUser)
}
