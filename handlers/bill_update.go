package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func BillUpdate(c *gin.Context) {
	var (
		reqBill *model.Bill
		userID  = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&reqBill); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpBill := service.GetBillByID(reqBill.BillID)
	if tmpBill == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}

	if tmpBill.UserID != userID {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	tmpBill.BillType = reqBill.BillType
	tmpBill.CategoryID = reqBill.CategoryID
	tmpBill.Fee = reqBill.Fee
	tmpBill.Remark = reqBill.Remark
	tmpBill.Year = reqBill.Year
	tmpBill.Mouth = reqBill.Mouth
	if err := service.UpdateBill(tmpBill); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
