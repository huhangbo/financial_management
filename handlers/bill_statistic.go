package handlers

import (
	"financial_management/consts"
	"financial_management/model"
	"financial_management/service"
	"financial_management/util"
	"github.com/gin-gonic/gin"
	"time"
)

type statistic struct {
	Fee          int            `json:"fee"`
	Month        int            `json:"month"`
	Day          int            `json:"day"`
	BillType     model.BillType `json:"bill_type"`
	BillTypeName string         `json:"bill_type_name"`
	CategoryName string         `json:"category_name"`
}

type statisticReq struct {
	IsAnalysis bool           `json:"is_analysis"`
	BillType   model.BillType `json:"bill_type"`
	Year       int            `json:"year"`
	Month      int            `json:"month"`
	BeginTime  time.Time      `json:"begin_time"`
	EndTime    time.Time      `json:"end_time"`
}

var (
	billTypeNameMap = map[model.BillType]string{
		model.Expend: "支出",
		model.Income: "收入",
	}
)

func BillStatistic(c *gin.Context) {
	var (
		req          *statisticReq
		userID       = c.GetInt(consts.UserID)
		list         []*statistic
		billTypeList []model.BillType
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

	switch {
	case req.Year != 0 && req.Month == 0:
		tmpBillList, err := service.GetBillByYear(billTypeList, userID, req.Year)
		if err != nil {
			util.Response(c, consts.SystemErrorCode, nil)
			return
		}
		var tmpStatistic []*statistic
		for i := 0; i < 24; i++ {
			tmpStatistic = append(tmpStatistic, &statistic{})
		}
		for i := 0; i < 12; i++ {
			tmpStatistic[i].BillTypeName, tmpStatistic[i+12].BillTypeName = "支出", "收入"
			tmpStatistic[i].Month, tmpStatistic[i+12].Month = i+1, i+1
			for _, item := range tmpBillList {
				if item.Month == i {
					if item.BillType == model.Expend {
						tmpStatistic[i].Fee += item.Fee
					} else {
						tmpStatistic[i+12].Fee += item.Fee
					}
				}
			}
		}
		list = tmpStatistic

	case !req.IsAnalysis && req.Year != 0 && req.Month != 0:
		tmpBillList, err := service.GetBillByMonth(billTypeList, userID, req.Year, req.Month)
		dayNum := util.GetYearMonthToDay(req.Year, req.Month)
		if err != nil {
			util.Response(c, consts.SystemErrorCode, nil)
			return
		}
		var tmpStatistic []*statistic
		for i := 0; i < dayNum*2; i++ {
			tmpStatistic = append(tmpStatistic, &statistic{})
		}
		for i := 0; i < dayNum; i++ {
			tmpStatistic[i].BillTypeName, tmpStatistic[i+dayNum].BillTypeName = billTypeNameMap[model.Expend], billTypeNameMap[model.Income]
			tmpStatistic[i].Day, tmpStatistic[i+dayNum].Day = i+1, i+1
			for _, item := range tmpBillList {
				if item.Day == i {
					if item.BillType == model.Expend {
						tmpStatistic[i].Fee += item.Fee
					} else {
						tmpStatistic[i+dayNum].Fee += item.Fee
					}
				}
			}
		}
		list = tmpStatistic

	case req.IsAnalysis:
		tmpBillList, err := service.GetBillAnalysis(billTypeList, userID, req.Year, req.Month)
		if err != nil {
			util.Response(c, consts.SystemErrorCode, nil)
			return
		}
		var (
			categoryIDs  []int
			categoryMap  = make(map[int]*model.Category)
			statisticMap = map[int]map[int]*statistic{
				req.Month:     {},
				req.Month - 1: {},
			}
			tmpList []*statistic
		)

		for _, bill := range tmpBillList {
			if bill.CategoryID == 0 {
				continue
			}
			categoryIDs = append(categoryIDs, bill.CategoryID)
		}
		if len(categoryIDs) != 0 {
			categoryMap, err = service.MGetCategory(categoryIDs)
			if err != nil {
				util.Response(c, consts.SystemErrorCode, nil)
				return
			}
		}
		for _, bill := range tmpBillList {
			if bill.Month == req.Month {
				if statisticMap[req.Month][bill.CategoryID] == nil {
					statisticMap[req.Month][bill.CategoryID] = &statistic{}
				}
				statisticMap[req.Month][bill.CategoryID].Fee += bill.Fee
				statisticMap[req.Month][bill.CategoryID].BillType = bill.BillType
			} else {
				if statisticMap[req.Month-1][bill.CategoryID] == nil {
					statisticMap[req.Month-1][bill.CategoryID] = &statistic{}
				}
				statisticMap[req.Month-1][bill.CategoryID].Fee += bill.Fee
				statisticMap[req.Month-1][bill.CategoryID].BillType = bill.BillType
			}

		}
		for month, tmpMap := range statisticMap {
			for categoryID, item := range tmpMap {
				tmpStatistic := &statistic{
					Fee:          item.Fee,
					BillType:     item.BillType,
					CategoryName: categoryMap[categoryID].CategoryDetail,
					BillTypeName: billTypeNameMap[item.BillType],
					Month:        month,
				}
				tmpList = append(tmpList, tmpStatistic)
			}
		}
		list = tmpList

	case req.Year == 0 && req.Month == 0:
		tmpBillList, err := service.GetBillByTime(billTypeList, userID, req.BeginTime, req.EndTime)
		if err != nil {
			util.Response(c, consts.SystemErrorCode, nil)
			return
		}
		var (
			categoryIDs  []int
			categoryMap  = make(map[int]*model.Category)
			statisticMap = make(map[int]*statistic)
			tmpList      []*statistic
		)
		for _, bill := range tmpBillList {
			if bill.CategoryID == 0 {
				continue
			}
			categoryIDs = append(categoryIDs, bill.CategoryID)
		}
		if len(categoryIDs) != 0 {
			categoryMap, err = service.MGetCategory(categoryIDs)
			if err != nil {
				util.Response(c, consts.SystemErrorCode, nil)
				return
			}
		}
		for _, bill := range tmpBillList {
			if statisticMap[bill.CategoryID] == nil {
				statisticMap[bill.CategoryID] = &statistic{}
			}
			statisticMap[bill.CategoryID].Fee += bill.Fee
			statisticMap[bill.CategoryID].BillType = bill.BillType
		}
		for categoryID, item := range statisticMap {
			tmpStatistic := &statistic{
				Fee:          item.Fee,
				BillType:     item.BillType,
				CategoryName: categoryMap[categoryID].CategoryDetail,
				BillTypeName: billTypeNameMap[item.BillType],
			}
			tmpList = append(tmpList, tmpStatistic)
		}
		list = tmpList

	default:
		util.Response(c, consts.ParamErrorCode, nil)
		return
	}

	util.Response(c, consts.SuccessCode, list)
	return
}
