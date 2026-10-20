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

func (s *UrlService) Encode(ctx *gin.Context, req dto.UrlReq) (*dto.UrlResp, error) {
	query := util.UrlEncodeQuery(req.Text)
	path := util.UrlEncodePath(req.Text)

	return &dto.UrlResp{
		Query: query,
		Path:  path,
	}, nil
}

func (s *UrlService) Decode(ctx *gin.Context, req dto.UrlReq) (*dto.UrlResp, error) {
	query, err := util.UrlDecodeQuery(req.Text)
	if err != nil {
		return nil, err
	}

	path, err := util.UrlDecodePath(req.Text)
	if err != nil {
		return nil, err
	}

	return &dto.UrlResp{
		Query: query,
		Path:  path,
	}, nil
}
