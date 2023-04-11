package main

import (
	"financial_management/setting"
	"github.com/gin-gonic/gin"
)

var (
	h *gin.Engine
)

func main() {
	gin.SetMode(setting.Config.Mode)
	h = gin.Default()
	setting.LoadConfig("./setting/config.json")

	initRouter()
	panic(h.Run())
}
