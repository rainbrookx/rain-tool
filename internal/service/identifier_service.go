package service

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/rain-tool/internal/dto"
	"github.com/rainbrookx/rain-tool/internal/util"
)

type IdentifierService struct {
}

func NewIdentifierService() *IdentifierService {
	return &IdentifierService{}
}

func (h *IdentifierService) Convert(ctx *gin.Context, req dto.IdentifierConvertReq) ([]*dto.IdentifierConvertResp, error) {
	var rs []*dto.IdentifierConvertResp

	originalTextList := req.OriginalTextList

	for _, originalText := range originalTextList {
		convert, err := util.IdentifierConvert(originalText)
		if err != nil {
			convert = &util.Identifier{}
		}

		resp := dto.IdentifierConvertResp{}
		resp.OriginalText = strings.TrimSpace(originalText)
		resp.Identifier = *convert

		rs = append(rs, &resp)
	}

	return rs, nil
}
