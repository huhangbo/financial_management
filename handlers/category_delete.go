package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/setting"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func CategoryDelete(c *gin.Context) {
	var (
		tmpCategory *model.Category
		userID      = c.GetInt(consts.UserID)
	)
	// 参数校验
	if err := c.BindJSON(&tmpCategory); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpCategory = service.GetCategoryByID(tmpCategory.CategoryID)
	if tmpCategory == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	// 非管理员操作无权限
	if userID != setting.Config.Admin.ID {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	if err := service.DeleteCategoryByID(tmpCategory.CategoryID); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
