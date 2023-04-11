package handlers

import (
	"financial_management/consts"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

func NotesDelete(c *gin.Context) {
	var (
		notesIDs []int
		userID   = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&notesIDs); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	notesList := service.GetNotesByIDs(notesIDs)
	if notesList == nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	for _, note := range notesList {
		if note.UserID != userID {
			util.Response(c, consts.PermissionErrorCode, nil)
			return
		}
	}
	if err := service.DeleteNotes(notesIDs); err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	util.Response(c, consts.SuccessCode, nil)
}
