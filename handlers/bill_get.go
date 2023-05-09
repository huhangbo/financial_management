package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
)

type BillGetReq struct {
	BillType  model.BillType `json:"bill_type"`
	Year      int            `json:"year"`
	Month     int            `json:"month"`
	BeginTime int            `json:"begin_time"`
	EndTime   int            `json:"end_time"`
}

type BillGetResp struct {
	BillList  []*model.Bill `json:"bill_list"`
	Total     int           `json:"total"`
	TotalPage int           `json:"total_page"`
}

func BillGet(c *gin.Context) {
	var (
		req          *BillGetReq
		billTypeList []model.BillType
		billList     []*model.Bill
		categoryIDs  []int
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

	if req.Month == 0 && req.Year == 0 {
		billList, err = service.GetUserBill(billTypeList, userID)
	} else {
		billList, err = service.GetBillByMonth(billTypeList, userID, req.Year, req.Month)
	}

	if err != nil {
		util.Response(c, consts.SystemErrorCode, nil)
		return
	}
	for _, bill := range billList {
		if bill.CategoryID == 0 {
			continue
		}
		categoryIDs = append(categoryIDs, bill.CategoryID)
	}
	if len(categoryIDs) != 0 {
		categoryMap, err := service.MGetCategory(categoryIDs)
		if err != nil {
			util.Response(c, consts.SystemErrorCode, nil)
			return
		}
		for _, bill := range billList {
			bill.Category = categoryMap[bill.CategoryID]
		}
	}
	//resp.BillList = billList
	//resp.Total = count
	//resp.TotalPage = (count-1)/req.Limit + 1
	util.Response(c, consts.SuccessCode, billList)
}
