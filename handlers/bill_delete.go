package handlers

import (
	"financial_management/consts"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func BillDelete(c *gin.Context) {
	var (
		billIDs []int
		userID  = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&billIDs); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	billList := service.GetBillByIDs(billIDs)
	if billList == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	for _, bill := range billList {
		if bill.UserID != userID {
			util.Response(c, consts.PermissionErrorCode, nil)
			return
		}
	}

	if err := service.DeleteBill(billIDs); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
