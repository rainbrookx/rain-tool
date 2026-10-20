package util

import "net/url"

// UrlEncodeQuery query 参数编码
func UrlEncodeQuery(s string) string {
	return url.QueryEscape(s)
}

// UrlDecodeQuery query 解码
func UrlDecodeQuery(s string) (string, error) {
	return url.QueryUnescape(s)
}

// UrlEncodePath url 路径编码
func UrlEncodePath(s string) string {
	return url.PathEscape(s)
}

// UrlDecodePath url 路径解码
func UrlDecodePath(s string) (string, error) {
	return url.PathUnescape(s)
}
