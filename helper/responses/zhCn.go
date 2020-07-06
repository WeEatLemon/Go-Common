package resp

var ZhCn = subMsgType{
	//基本操作
	Success:           "成功",
	UnknownError:      "未知错误，请联系支持",
	ParameterError:    "参数错误",
	OperationFailed:   "操作失败",
	MissingParameters: "缺少参数",
	InvalidOperation:  "无效操作",
	UserNotAuthorize:  "用户未授权",
	FrequencyTooFast:  "频率过快",
	DataAlreadyExists: "数据已存在",
	DataDoesNotExist:  "数据不存在",

	//用户
	UserExists:    "用户已存在",
	UserNotExists: "用户不存在",
	LackBalance:   "余额不足",

	//验证
	PasswordError:   "用户名或密码错误",
	TokenExpire:     "token过期",
	GoogleCodeError: "谷歌验证码错误",
}
