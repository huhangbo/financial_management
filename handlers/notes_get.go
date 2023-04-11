package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

type NotesGetReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type NotesGetResp struct {
	NotesList  []*model.Notes `json:"notes_list"`
	HasMore    bool           `json:"has_more"`
	NextOffset int            `json:"next_offset"`
}

func NotesGet(c *gin.Context) {
	var (
		req  *NotesGetReq
		resp = &NotesGetResp{}
	)
	if err := c.BindJSON(&req); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}
	notesList := service.GetNotesList(c.GetInt(consts.UserID), req.Offset, req.Limit+1)
	if notesList == nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	if len(notesList) > req.Limit {
		resp.HasMore = true
		notesList = notesList[:req.Limit]
	}
	resp.NextOffset = req.Offset + req.Limit
	resp.NotesList = notesList
	util.Response(c, consts.SuccessCode, resp)
}
