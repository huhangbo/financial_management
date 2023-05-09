package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func BudgetAdd(c *gin.Context) {
	var (
		reqBudget *model.Budget
	)
	if err := c.BindJSON(&reqBudget); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	reqBudget.UserID = c.GetInt(consts.UserID)
	if err := service.AddBudget(reqBudget); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	respBudget := service.GetBudgetByUserID(c.GetInt(consts.UserID), reqBudget.Year, reqBudget.Month)
	if reqBudget == nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, respBudget)
}
