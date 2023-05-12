package handlers

import (
	"financial_management/consts"
	"financial_management/service"
	"financial_management/setting"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func CategoryDelete(c *gin.Context) {
	var (
		categoryIDs []int
		userID      = c.GetInt(consts.UserID)
	)
	// 参数校验
	if err := c.BindJSON(&categoryIDs); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	// 非管理员操作无权限
	if userID != setting.Config.Admin.ID {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}

	_, err := service.MGetCategory(categoryIDs)
	if err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}

	if err := service.DeleteCategory(categoryIDs); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
