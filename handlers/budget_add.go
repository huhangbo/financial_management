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
		tmpBudget *model.Budget
	)
	if err := c.BindJSON(&tmpBudget); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpBudget.UserID = c.GetInt(consts.UserID)
	if err := service.AddBudget(tmpBudget); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
