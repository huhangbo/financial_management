package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func NewsAdd(c *gin.Context) {
	var (
		tmpNews *model.News
	)
	if err := c.BindJSON(&tmpNews); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}

	if err := service.AddNews(tmpNews); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
