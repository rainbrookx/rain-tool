package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rainbrookx/rain-tool/internal/dto"
	"github.com/rainbrookx/rain-tool/internal/util"
)

type UrlService struct {
}

func NewUrlService() *UrlService {
	return &UrlService{}
}

func (s *UrlService) EncodeQuery(ctx *gin.Context, req *dto.UrlReq) (*dto.UrlResp, error) {
	return &dto.UrlResp{Text: util.UrlEncodeQuery(req.Text)}, nil
}

func (s *UrlService) DecodeQuery(ctx *gin.Context, req *dto.UrlReq) (*dto.UrlResp, error) {
	data, err := util.UrlDecodeQuery(req.Text)
	if err != nil {
		return nil, err
	}

	return &dto.UrlResp{Text: data}, nil
}

func (s *UrlService) EncodePath(ctx *gin.Context, req *dto.UrlReq) (*dto.UrlResp, error) {
	return &dto.UrlResp{Text: util.UrlEncodePath(req.Text)}, nil
}

func (s *UrlService) DecodePath(ctx *gin.Context, req *dto.UrlReq) (*dto.UrlResp, error) {
	data, err := util.UrlDecodePath(req.Text)
	if err != nil {
		return nil, err
	}

	return &dto.UrlResp{Text: data}, nil
}
