package dto

// Result 统一响应体
type Result struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"` // omitempty data为空不返回
}

// Success 成功响应 A1000
func Success(data any) *Result {
	return &Result{
		Code: "A1000",
		Msg:  "成功",
		Data: data,
	}
}

// Fail 失败响应 A2000
func Fail(msg string) *Result {
	return &Result{
		Code: "A2000",
		Msg:  msg,
	}
}
