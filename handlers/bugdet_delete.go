package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func BudgetDelete(c *gin.Context) {
	var (
		reqBudget *model.Budget
	)
	if err := c.BindJSON(&reqBudget); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpBudget := service.GetBudgetByID(reqBudget.BudgetId)
	if tmpBudget == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	if tmpBudget.UserID != c.GetInt(consts.UserID) {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	if err := service.DeleteBudget(reqBudget.BudgetId); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
