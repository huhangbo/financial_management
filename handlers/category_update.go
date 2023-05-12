package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/setting"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func CategoryUpdate(c *gin.Context) {
	var (
		reqCategory *model.Category
		userID      = c.GetInt(consts.UserID)
	)
	// 参数校验
	if err := c.BindJSON(&reqCategory); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpCategory := service.GetCategoryByID(reqCategory.CategoryID) // 先查询出旧的
	if tmpCategory == nil {
		util.Response(c, consts.ParamErrorCode, nil)
	}
	// 非管理员操作无权限
	if userID != setting.Config.Admin.ID {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	tmpCategory.CategoryDetail = reqCategory.CategoryDetail
	if err := service.UpdateCategoryInfo(tmpCategory); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
