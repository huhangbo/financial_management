package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func NotesUpdate(c *gin.Context) {
	var (
		reqNotes *model.Notes
	)
	if err := c.BindJSON(&reqNotes); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	tmpNotes := service.GetNotesByID(reqNotes.NotesID)
	if tmpNotes == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	if tmpNotes.UserID != c.GetInt(consts.UserID) {
		util.Response(c, consts.PermissionErrorCode, nil)
		return
	}
	tmpNotes.Title = reqNotes.Title
	tmpNotes.Detail = reqNotes.Detail
	if err := service.UpdateNotes(tmpNotes); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
