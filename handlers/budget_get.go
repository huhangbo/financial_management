package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func BudgetGet(c *gin.Context) {
	var (
		reqBudget *model.Budget
	)
	if err := c.BindJSON(&reqBudget); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpBudget := service.GetBudgetByUserID(c.GetInt(consts.UserID), reqBudget.Year, reqBudget.Mouth)
	if tmpBudget == nil {
		util.Response(c, consts.EmptyCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, tmpBudget)
}
