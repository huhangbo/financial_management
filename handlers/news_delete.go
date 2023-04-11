package handlers

import (
	"financial_management/consts"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func NewsDelete(c *gin.Context) {
	var (
		newsIDs []int
		userID  = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&newsIDs); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	newsList := service.GetNewsByIDs(newsIDs)
	if newsList == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	for _, news := range newsList {
		if news.UserID != userID {
			util.Response(c, consts.PermissionErrorCode, nil)
			return
		}
	}
	if err := service.DeleteNews(newsIDs); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
