package handler

import "github.com/gin-gonic/gin"

type UrlHandler struct {
}

func NewUrlHandler() *UrlHandler {
	return &UrlHandler{}
}

// EncodeQuery query 参数编码
func (h *UrlHandler) EncodeQuery(ctx *gin.Context) {

}

// DecodeQuery query 解码
func (h *UrlHandler) DecodeQuery(ctx *gin.Context) {

}

// EncodePath url 路径编码
func (h *UrlHandler) EncodePath(ctx *gin.Context) {

}

// DecodePath url 路径解码
func (h *UrlHandler) DecodePath(ctx *gin.Context) {

}
