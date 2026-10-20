package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/rain-tool/internal/router"
)

func main() {
	initMain()

	engine := gin.Default()

	group := engine.Group("/api")
	router.InitRouter(group)

	if err := engine.Run("127.0.0.1:8080"); err != nil {
		slog.Info(err.Error())
		return
	}

	// TODO 可以考虑加上优雅退出
}

func initMain() {

}
