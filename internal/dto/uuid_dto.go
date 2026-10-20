package dto

type UuidResp struct {
	V4         string `json:"v4"`
	V7         string `json:"v7"`
	V4NoHyphen string `json:"v4NoHyphen"`
	V7NoHyphen string `json:"v7NoHyphen"`
}
