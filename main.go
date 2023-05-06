package main

import (
	"financial_management/setting"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	h *gin.Engine
)

func main() {
	setting.LoadConfig("./setting/config.json")

	//f, _ := os.Create(setting.Config.LogConfig.Path)
	//gin.DefaultWriter = io.MultiWriter(f)

	gin.SetMode(setting.Config.Mode)

	h = gin.Default()

	initRouter()

	panic(h.Run(fmt.Sprintf(":%d", setting.Config.Port)))
}
