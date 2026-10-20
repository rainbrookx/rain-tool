package handler

import "github.com/rainbrookx/rain-tool/internal/service"

type AllHandler struct {
	Url        *UrlHandler
	Identifier *IdentifierHandler
	Uuid       *UuidHandler
}

func NewAllHandler(allService *service.AllService) *AllHandler {
	var h AllHandler
	h.Url = NewUrlHandler(allService)
	h.Identifier = NewIdentifierHandler(allService)
	h.Uuid = NewUuidHandler(allService)

	return &h
}
