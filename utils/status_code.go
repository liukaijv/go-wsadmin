package utils

const StatusCodeMustLogin = 401
const StatusCodeMustReLogin = 402
const StatusCodeForbidden = 403
const StatusCodeNotFound = 404
const StatusCodeParamError = 405

var StatusCode map[int]string = map[int]string{
	StatusCodeMustLogin:   "未登录",
	StatusCodeMustReLogin: "需要重新登录",
	StatusCodeForbidden:   "访问被禁止",
	StatusCodeNotFound:    "未找到",
	StatusCodeParamError:  "参数错误",
}

func GetStatusDisplay(code int) string {
	if val, ok := StatusCode[code]; ok {
		return val
	}
	return ""
}
