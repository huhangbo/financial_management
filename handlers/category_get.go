package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CategoryGetList(c *gin.Context) {
	var (
		categoryList []*model.Category
		categoryType = c.Query("type")
	)
	typeNum, _ := strconv.ParseInt(categoryType, 10, 64)
	categoryList, err := service.GetCategoryList(model.BillType(typeNum))
	if err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, categoryList)
}
