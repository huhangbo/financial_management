package main

import (
	"financial_management/setting"
	"github.com/gin-gonic/gin"
)

var (
	h *gin.Engine
)

func main() {
	setting.LoadConfig("./setting/config.json")
	gin.SetMode(setting.Config.Mode)
	h = gin.Default()

	initRouter()
	panic(h.Run())
}
