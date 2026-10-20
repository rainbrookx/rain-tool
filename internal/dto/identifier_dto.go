package dto

import "github.com/rainbrookx/rain-tool/internal/util"

type IdentifierConvertReq struct {
	OriginalTextList []string `json:"originalTextList"`
}

type IdentifierConvertResp struct {
	OriginalText string          `json:"originalText"`
	Identifier   util.Identifier `json:"identifier"`
}

type IdentifierPrefixAndSuffixReq struct {
	OriginalTextList    []string               `json:"originalTextList"`
	PrefixAndSuffixList []util.PrefixAndSuffix `json:"prefixAndSuffixList"`
}

type IdentifierPrefixAndSuffixResp struct {
	OriginalText string   `json:"originalText"`
	ModifiedList []string `json:"modifiedList"`
}
