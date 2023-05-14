package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/setting"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func CategoryAdd(c *gin.Context) {
	var (
		reqCategory *model.Category
		userID      = c.GetInt(consts.UserID)
	)

	if err := c.BindJSON(&reqCategory); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	if userID != setting.Config.Admin.ID {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	if err := service.AddCategory(reqCategory); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	lastCategory := service.GetLastCategory()
	if lastCategory == nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, lastCategory)
}
