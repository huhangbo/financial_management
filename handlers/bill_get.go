package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

type BillGetReq struct {
	BillType model.BillType `json:"bill_type"`
	Year     int            `json:"year"`
	Mouth    int            `json:"mouth"`
	Page     int            `json:"page"`
	Limit    int            `json:"limit"`
}

type BillGetResp struct {
	BillList  []*model.Bill `json:"bill_list"`
	Total     int           `json:"total"`
	TotalPage int           `json:"total_page"`
}

func BillGet(c *gin.Context) {
	var (
		req          *BillGetReq
		resp         = &BillGetResp{}
		billTypeList []model.BillType
		billList     []*model.Bill
		categoryIDs  []int
		word         = c.Query("word")
		err          error
		userID       = c.GetInt(consts.UserID)
	)
	if err := c.BindJSON(&req); err != nil {
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}

	if req.BillType == 0 {
		billTypeList = []model.BillType{model.Expend, model.Income}
	} else {
		billTypeList = append(billTypeList, req.BillType)
	}
	if len(word) == 0 {
		billList, err = service.GetBillByTime(billTypeList, userID, req.Year, req.Mouth, req.Page, req.Limit)
	} else {
		billList, err = service.SearchBillByTime(billTypeList, word, userID, req.Year, req.Mouth, req.Page, req.Limit)
	}

	count := service.GetBillCount(billTypeList, userID, req.Year, req.Mouth)
	if err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
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
	resp.Total = count
	resp.TotalPage = (count-1)/req.Limit + 1
	util.Response(c, consts.SuccessCode, resp)
}
