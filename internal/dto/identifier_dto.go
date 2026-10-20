package dto

import "github.com/rainbrookx/rain-tool/internal/util"

type IdentifierConvertReq struct {
	OriginalTextList []string `json:"originalTextList"`
}

type IdentifierConvertResp struct {
	OriginalText string          `json:"originalText"`
	Identifier   util.Identifier `json:"identifier"`
}
