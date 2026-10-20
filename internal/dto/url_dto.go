package dto

type UrlReq struct {
	Text string `json:"text"`
}

type UrlResp struct {
	Query string `json:"query"`
	Path  string `json:"path"`
}
