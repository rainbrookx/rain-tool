package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/rain-tool/internal/dto"
	"github.com/rainbrookx/rain-tool/internal/service"
	"github.com/rainbrookx/rain-tool/internal/util_inner"
)

type UuidHandler struct {
}

var uuidService *service.UuidService

func NewUuidHandler(allService *service.AllService) *UuidHandler {
	uuidService = allService.Uuid

	return &UuidHandler{}
}

func (h *UuidHandler) Generate(ctx *gin.Context) {
	var req dto.UuidReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(200, dto.Fail("JSON 解析失败"))
		return
	}

	req.Quantity = util_inner.ClampNumber(req.Quantity, 1, 1000)

	data, err := uuidService.Generate(ctx, req)
	if err != nil {
		ctx.JSON(200, dto.Fail("Generate 失败"))
		return
	}

	ctx.JSON(200, dto.Success(data))
}
