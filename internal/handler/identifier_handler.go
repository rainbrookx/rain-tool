package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/rain-tool/internal/dto"
	"github.com/rainbrookx/rain-tool/internal/service"
)

type IdentifierHandler struct {
}

var identifierService *service.IdentifierService

func NewIdentifierHandler(allService *service.AllService) *IdentifierHandler {
	identifierService = allService.Identifier
	return &IdentifierHandler{}
}

func (h *IdentifierHandler) Convert(ctx *gin.Context) {
	var req dto.IdentifierConvertReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(200, dto.Fail("JSON 解析失败"))
		return
	}

	data, err := identifierService.Convert(ctx, req)
	if err != nil {
		ctx.JSON(200, dto.Fail("Convert 失败"))
		return
	}

	ctx.JSON(200, dto.Success(data))
}

func (h *IdentifierHandler) PrefixAndSuffix(ctx *gin.Context) {
	var req dto.IdentifierPrefixAndSuffixReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(200, dto.Fail("JSON 解析失败"))
		return
	}

	data, err := identifierService.PrefixAndSuffix(ctx, req)
	if err != nil {
		ctx.JSON(200, dto.Fail("PrefixAndSuffix 失败"))
		return
	}

	ctx.JSON(200, dto.Success(data))
}
