package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/rain-tool/internal/handler"
	"github.com/rainbrookx/rain-tool/internal/infrastructure/config"
	"github.com/rainbrookx/rain-tool/internal/router"
	"github.com/rainbrookx/rain-tool/internal/service"
)

func main() {
	initMain()

	// 依赖注入
	allService := service.NewAllService()
	allHandler := handler.NewAllHandler(allService)

	// gin 初始化和注册路由
	engine := gin.Default()
	router.RegisterRoute(engine, allHandler)

	if err := engine.Run(config.C.RainTool.Addr); err != nil {
		slog.Info(err.Error())
		return
	}

	// TODO 可以考虑加上优雅退出
}

func initMain() {
	// TODO 将config、DB、Redis等以返回值形式给 main，不要使用全局变量导出
	config.InitConfig()
}
