package service

import (
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
	uuid47 := util.UuidGenerate()

	return &dto.UuidResp{
		V4:         uuid47.V4,
		V7:         uuid47.V7,
		V4NoHyphen: uuid47.V4NoHyphen,
		V7NoHyphen: uuid47.V7NoHyphen,
	}, nil
}
