package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
	"time"
)

type BillGetReq struct {
	BillType   model.BillType `json:"bill_type"`
	BeginYear  int            `json:"begin_year"`
	BeginMouth int            `json:"begin_mouth"`
	BeginDay   int            `json:"begin_day"`
	EndYear    int            `json:"end_year"`
	EndMouth   int            `json:"end_mouth"`
	EndDay     int            `json:"end_day"`
	Offset     int            `json:"offset"`
	Limit      int            `json:"limit"`
}

type BillGetResp struct {
	BillList   []*model.Bill `json:"bill_list"`
	NextOffset int           `json:"next_offset"`
	HasMore    bool          `json:"has_more"`
}

func BillGet(c *gin.Context) {
	var (
		req          *BillGetReq
		resp         = &BillGetResp{}
		billTypeList []model.BillType
		categoryIDs  []int
	)
	if err := c.BindJSON(&req); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}

	beginTime := time.Date(req.BeginYear, time.Month(req.BeginMouth), req.BeginDay, 0, 0, 0, 0, nil)
	endTime := time.Date(req.EndYear, time.Month(req.EndMouth), req.EndDay, 0, 0, 0, 0, nil)
	if req.BillType == 0 {
		billTypeList = []model.BillType{model.Expend, model.Income}
	} else {
		billTypeList = append(billTypeList, req.BillType)
	}
	billList, err := service.GetBillByTime(billTypeList, beginTime, endTime, req.Limit+1, req.Offset)
	if err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	if len(billList) > req.Limit {
		resp.HasMore = true
		billList = billList[:req.Limit]
	}
	for _, bill := range billList {
		categoryIDs = append(categoryIDs, bill.CategoryID)
	}
	categoryMap, err := service.MGetCategory(categoryIDs)
	if err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	for _, bill := range billList {
		bill.Category = categoryMap[bill.CategoryID]
	}
	resp.BillList = billList
	resp.NextOffset = req.Offset + req.Limit
	util.Response(c, consts.SuccessCode, resp)
}
