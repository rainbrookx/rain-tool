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

// EncodeQuery query 参数编码
func (h *UrlHandler) EncodeQuery(ctx *gin.Context) {
	var req dto.UrlReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(200, dto.Fail("JSON 解析失败"))
		return
	}

	data, err := urlService.EncodeQuery(ctx, req)
	if err != nil {
		ctx.JSON(200, dto.Fail("EncodeQuery 失败"))
		return
	}

	ctx.JSON(200, dto.Success(data))
}

// DecodeQuery query 解码
func (h *UrlHandler) DecodeQuery(ctx *gin.Context) {
	var req dto.UrlReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(200, dto.Fail("JSON 解析失败"))
		return
	}

	data, err := urlService.DecodeQuery(ctx, req)
	if err != nil {
		ctx.JSON(200, dto.Fail("DecodeQuery 失败"))
		return
	}

	ctx.JSON(200, dto.Success(data))
}

// EncodePath url 路径编码
func (h *UrlHandler) EncodePath(ctx *gin.Context) {
	var req dto.UrlReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(200, dto.Fail("JSON 解析失败"))
		return
	}

	data, err := urlService.EncodePath(ctx, req)
	if err != nil {
		ctx.JSON(200, dto.Fail("EncodePath 失败"))
		return
	}

	ctx.JSON(200, dto.Success(data))
}

// DecodePath url 路径解码
func (h *UrlHandler) DecodePath(ctx *gin.Context) {
	var req dto.UrlReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(200, dto.Fail("JSON 解析失败"))
		return
	}

	data, err := urlService.DecodePath(ctx, req)
	if err != nil {
		ctx.JSON(200, dto.Fail("DecodePath 失败"))
		return
	}

	ctx.JSON(200, dto.Success(data))
}
