package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/setting"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func NewsUpdate(c *gin.Context) {
	var (
		reqNews *model.News
		userID  = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&reqNews); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	if userID != setting.Config.Admin.ID {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	tmpNews := service.GetNewsByID(reqNews.NewsID)
	if tmpNews == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpNews.Title = reqNews.Title
	tmpNews.Detail = reqNews.Detail
	tmpNews.Source = reqNews.Source
	if err := service.UpdateNews(tmpNews); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
