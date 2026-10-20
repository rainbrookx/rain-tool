package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/rain-tool/internal/dto"
	"github.com/rainbrookx/rain-tool/internal/service"
)

type UrlHandler struct {
}

var urlService *service.UrlService

func NewUrlHandler(allService *service.AllService) *UrlHandler {
	urlService = allService.Url

	return &UrlHandler{}
}

func (h *UrlHandler) Encode(ctx *gin.Context) {
	var req dto.UrlReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(200, dto.Fail("JSON 解析失败"))
		return
	}

	data, err := urlService.Encode(ctx, req)
	if err != nil {
		ctx.JSON(200, dto.Fail("Encode 失败"))
		return
	}

	ctx.JSON(200, dto.Success(data))
}

func (h *UrlHandler) Decode(ctx *gin.Context) {
	var req dto.UrlReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(200, dto.Fail("JSON 解析失败"))
		return
	}

	data, err := urlService.Decode(ctx, req)
	if err != nil {
		ctx.JSON(200, dto.Fail("Decode 失败"))
		return
	}

	ctx.JSON(200, dto.Success(data))
}
