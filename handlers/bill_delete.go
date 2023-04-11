package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func BillDelete(c *gin.Context) {
	var (
		tmpBill *model.Bill
		userID  = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&tmpBill); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpBill = service.GetBillByID(tmpBill.BillID)
	if tmpBill == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	if tmpBill.UserID != userID {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	if err := service.DeleteBIll(tmpBill.BillID); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
