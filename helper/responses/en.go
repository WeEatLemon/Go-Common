package resp

var En = subMsgType{
	//基本操作
	Success:           "success",
	UnknownError:      "Unknown error, please contact support",
	ParameterError:    "Parameter error",
	OperationFailed:   "Operation failed",
	MissingParameters: "Missing parameters",
	InvalidOperation:  "Invalid operation",
	UserNotAuthorize:  "User is not authorized",
	FrequencyTooFast:  "Frequency too fast",
	DataAlreadyExists: "Data already exists",
	DataDoesNotExist:  "Data does not exist",

	//用户
	UserExists:    "User already exists",
	UserNotExists: "User does not exist",
	LackBalance:   "Insufficient balance",

	//验证
	PasswordError:   "Username or Password error",
	TokenExpire:     "Token expire",
	GoogleCodeError: "Google code error",
}
