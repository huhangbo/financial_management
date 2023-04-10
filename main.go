package financial_management

import (
	"financial_management/setting"
	"github.com/cloudwego/hertz/pkg/app/server"
)

var (
	h *server.Hertz
)

func main() {
	h = server.Default()
	initRouter()
	config := setting.LoadConfig("./setting/config.json")
	setting.InitMysql(config.SqlConfig)
	h.Spin()
}
