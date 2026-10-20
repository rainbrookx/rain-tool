package util

import "net/url"

// EncodeQuery query 参数编码
func EncodeQuery(s string) string {
	return url.QueryEscape(s)
}

// DecodeQuery query 解码
func DecodeQuery(s string) (string, error) {
	return url.QueryUnescape(s)
}

// EncodePath url 路径编码
func EncodePath(s string) string {
	return url.PathEscape(s)
}

// DecodePath url 路径解码
func DecodePath(s string) (string, error) {
	return url.PathUnescape(s)
}
