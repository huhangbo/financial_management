package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func CategoryGetList(c *gin.Context) {
	var (
		categoryList []*model.Category
	)
	categoryList, err := service.GetCategoryList()
	if err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, categoryList)
}
