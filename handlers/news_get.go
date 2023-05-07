package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

type NewsGetResp struct {
	NewsList []*model.News `json:"news_list"`
}

func NewsGet(c *gin.Context) {
	var (
		resp = &NewsGetResp{}
	)
	newsList := service.GetNewsList()
	if newsList == nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	resp.NewsList = newsList
	util.Response(c, consts.SuccessCode, resp)
}
