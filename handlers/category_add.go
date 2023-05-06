package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
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
	reqCategory.UserID = userID
	if err := service.AddCategory(reqCategory); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
