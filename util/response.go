package util

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"financial_management/consts"
)

func Response(ctx *gin.Context, code int, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": consts.ResponseCode2Msg,
		"data":    data,
	})
}
