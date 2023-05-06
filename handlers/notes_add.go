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
		userID   = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&tmpNotes); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpNotes.UserID = userID
	if err := service.AddNotes(tmpNotes); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
