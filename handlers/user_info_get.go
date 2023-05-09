package handlers

import (
	"financial_management/consts"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func UserInfoGet(c *gin.Context) {
	var (
		userID = c.GetInt(consts.UserID)
	)
	userInfo := service.GetUserByID(userID)
	if userInfo == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, UserLoginResp{
		User: userInfo,
	})
}
