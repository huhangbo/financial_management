package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
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
	tmpUser := service.GetUserByID(uid)
	if tmpUser == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpUser.UserName = reqUser.UserName
	tmpUser.Telephone = reqUser.Telephone
	if err := service.UpdateUser(tmpUser); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
