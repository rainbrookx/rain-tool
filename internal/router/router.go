package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/rain-tool/internal/handler"
)

func RegisterRoute(engine *gin.Engine, allHandler *handler.AllHandler) {
	group := engine.Group("/api")

	bindUrlHandler(group, "/url", allHandler.Url)
}

func bindUrlHandler(group *gin.RouterGroup, relativePath string, handler_ *handler.UrlHandler) {
	r := group.Group(relativePath)

	r.POST("/encode-query", handler_.EncodeQuery)
	r.POST("/decode-query", handler_.DecodeQuery)
	r.POST("/encode-path", handler_.EncodePath)
	r.POST("/decode-path", handler_.DecodePath)
}
