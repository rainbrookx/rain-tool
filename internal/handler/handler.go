package handler

import "github.com/rainbrookx/rain-tool/internal/service"

type AllHandler struct {
	Url *UrlHandler
}

func NewAllHandler(allService *service.AllService) *AllHandler {
	var h AllHandler
	h.Url = NewUrlHandler(allService)

	return &h
}
