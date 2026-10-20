package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/rain-tool/internal/handler"
)

func InitRouter(group *gin.RouterGroup) {
	bindUrlHandler(group, "/url")
}

func bindUrlHandler(group *gin.RouterGroup, relativePath string) {
	r := group.Group(relativePath)
	h := handler.NewUrlHandler()

	r.POST("/encode-query", h.EncodeQuery)
	r.POST("/decode-query", h.DecodeQuery)
	r.POST("/encode-path", h.EncodePath)
	r.POST("/decode-path", h.DecodePath)
}
