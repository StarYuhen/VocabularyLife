package server

// 请求顶层状态码

const (
	MissingParameters = -1 // 缺少参数
	InternalError     = -2 // 内部错误
	ParameterError    = -3 // 参数错误
	UnknownError      = -4 // 未知错误
	NotOwned          = -5 // 参数在内部未拥有或者不存在
)

// Return 返回值结果
type Return struct {
	Code int         `json:"code"` // 响应标识码
	Data interface{} `json:"data"` // 响应数据
	Msg  string      `json:"msg"`  // 请求返回含义
}

// Success 请求成功
func Success(data interface{}, msg string) Return {
	return Return{
		Code: 200,
		Data: data,
		Msg:  msg,
	}
}

// MissingParametersFun 请求缺少参数
func MissingParametersFun(msg string) Return {
	return Return{
		Code: MissingParameters,
		Data: nil,
		Msg:  msg,
	}
}

// InternalErrorFun 内部错误
func InternalErrorFun(msg string) Return {
	return Return{
		Code: InternalError,
		Data: nil,
		Msg:  msg,
	}
}

// ParameterErrorFun 参数错误
func ParameterErrorFun(msg string) Return {
	return Return{
		Code: ParameterError,
		Data: nil,
		Msg:  msg,
	}
}

// UnknownErrorFun 未知错误
func UnknownErrorFun(msg string) Return {
	return Return{
		Code: UnknownError,
		Data: nil,
		Msg:  msg,
	}
}

// NotOwnedFun 参数内部不存在
func NotOwnedFun(msg string) Return {
	return Return{
		Code: NotOwned,
		Data: nil,
		Msg:  msg,
	}
}
