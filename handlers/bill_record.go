package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func BillRecord(c *gin.Context) {
	var (
		reqBill *model.Bill
		userID  = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&reqBill); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	categoryInfo := service.GetCategoryByID(reqBill.CategoryID)
	if categoryInfo == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	reqBill.Category = categoryInfo
	reqBill.UserID = userID
	if err := service.AddBill(reqBill); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
