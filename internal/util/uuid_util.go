package util

import (
	"strings"
	"uuid"
)

type Uuid47 struct {
	V4         string `json:"v4"`
	V7         string `json:"v7"`
	V4NoHyphen string `json:"v4NoHyphen"`
	V7NoHyphen string `json:"v7NoHyphen"`
}

func UuidGenerate() *Uuid47 {
	v4 := uuid.NewV4().String()
	v7 := uuid.NewV7().String()

	return &Uuid47{
		V4:         v4,
		V7:         v7,
		V4NoHyphen: strings.ReplaceAll(v4, "-", ""),
		V7NoHyphen: strings.ReplaceAll(v7, "-", ""),
	}
}
