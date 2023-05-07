package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

type NotesGetResp struct {
	NotesList []*model.Notes `json:"notes_list"`
}

func NotesGet(c *gin.Context) {
	var (
		resp   = &NotesGetResp{}
		userID = c.GetInt(consts.UserID)
	)
	notesList := service.GetNotesList(userID)
	if notesList == nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	resp.NotesList = notesList
	util.Response(c, consts.SuccessCode, resp)
}
