package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/rain-tool/internal/dto"
	"github.com/rainbrookx/rain-tool/internal/service"
)

type UuidHandler struct {
}

var uuidService *service.UuidService

func NewUuidHandler(allService *service.AllService) *UuidHandler {
	uuidService = allService.Uuid

	return &UuidHandler{}
}

func (h *UuidHandler) Generate(ctx *gin.Context) {
	data, err := uuidService.Generate(ctx)
	if err != nil {
		ctx.JSON(200, dto.Fail("Generate 失败"))
		return
	}

	ctx.JSON(200, dto.Success(data))
}
