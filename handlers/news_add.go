package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/setting"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func NewsAdd(c *gin.Context) {
	var (
		tmpNews *model.News
		userID  = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&tmpNews); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	if userID != setting.Config.Admin.ID {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	if err := service.AddNews(tmpNews); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	respNews := service.GetLastNews()
	if respNews == nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, respNews)
}
