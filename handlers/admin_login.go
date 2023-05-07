package handlers

import (
	"financial_management/consts"
	"financial_management/middleware"
	"financial_management/model"
	"financial_management/setting"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	var (
		reqUser *model.User
	)
	if err := c.BindJSON(&reqUser); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	if reqUser.Telephone != setting.Config.Admin.ID || reqUser.Password != setting.Config.Admin.Password {
		util.Response(c, consts.PasswordErrorCode, nil)
		return
	}
	token := middleware.GenerateToken(reqUser.Telephone)
	util.Response(c, consts.SuccessCode, gin.H{
		"token": token,
	})
}
