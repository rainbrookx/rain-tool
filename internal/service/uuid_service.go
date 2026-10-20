package service

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/rain-tool/internal/dto"
	"github.com/rainbrookx/rain-tool/internal/util"
)

type UuidService struct {
}

func NewUuidService() *UuidService {
	return &UuidService{}
}

func (s *UuidService) Generate(ctx *gin.Context) (*dto.UuidResp, error) {
	v4 := util.UuidGenerateV4()
	v7 := util.UuidGenerateV7()

	return &dto.UuidResp{
		V4:         v4,
		V7:         v7,
		V4NoHyphen: strings.ReplaceAll(v4, "-", ""),
		V7NoHyphen: strings.ReplaceAll(v7, "-", ""),
	}, nil
}
