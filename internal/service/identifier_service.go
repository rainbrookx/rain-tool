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

func (s *IdentifierService) Convert(ctx *gin.Context, req dto.IdentifierConvertReq) ([]*dto.IdentifierConvertResp, error) {
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

func (s *IdentifierService) PrefixAndSuffix(
	ctx *gin.Context, req dto.IdentifierPrefixAndSuffixReq,
) ([]*dto.IdentifierPrefixAndSuffixResp, error) {
	var rs []*dto.IdentifierPrefixAndSuffixResp

	originalTextList := req.OriginalTextList

	for _, originalText := range originalTextList {
		ps := util.IdentifierPrefixAndSuffix(originalText, req.PrefixAndSuffixList)

		resp := dto.IdentifierPrefixAndSuffixResp{}
		resp.OriginalText = originalText
		resp.ModifiedList = ps

		rs = append(rs, &resp)
	}

	return rs, nil
}
