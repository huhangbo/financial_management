package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

type NewsGetReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type NewsGetResp struct {
	NewsList   []*model.News `json:"notes_list"`
	HasMore    bool          `json:"has_more"`
	NextOffset int           `json:"next_offset"`
}

func NewsGet(c *gin.Context) {
	var (
		req  *NewsGetReq
		resp = &NewsGetResp{}
	)
	if err := c.BindJSON(&req); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	newsList := service.GetNewsList(req.Offset, req.Limit+1)
	if newsList == nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	if len(newsList) > req.Limit {
		resp.HasMore = true
		newsList = newsList[:req.Limit]
	}
	resp.NextOffset = req.Offset + req.Limit
	resp.NewsList = newsList
	util.Response(c, consts.SuccessCode, resp)
}
