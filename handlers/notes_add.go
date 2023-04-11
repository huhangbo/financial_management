package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func NotesAdd(c *gin.Context) {
	var (
		tmpNotes *model.Notes
	)
	if err := c.BindJSON(&tmpNotes); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}

	if err := service.AddNotes(tmpNotes); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
