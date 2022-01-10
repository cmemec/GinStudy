package main

import (
	controller "GinStudy/Controller"
	tool "GinStudy/Tool"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig("./Config/app.json")
	if err != nil {
		panic(err)
	}
	engine := gin.Default()
	registerRouter(engine)
	tool.OrmEngine(cfg)
	engine.Run(cfg.AppHost + ":" + cfg.AppPort)
}

// 路由
func registerRouter(engine *gin.Engine) {
	new(controller.SendMsg).Router(engine)
}
