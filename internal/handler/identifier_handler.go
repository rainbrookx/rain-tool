package handler

import "github.com/rainbrookx/rain-tool/internal/service"

type IdentifierHandler struct {
}

func NewIdentifierHandler(allService *service.AllService) *IdentifierHandler {
	return &IdentifierHandler{}
}
