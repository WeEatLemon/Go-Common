package resp

import (
	"github.com/WeEatLemon/Go-Common/language"
	"github.com/gin-gonic/gin"
	"net/http"
)

var RESP *Resp

type Resp struct {
	lan string
}

func NewResp() *Resp {
	return RESP
}

func InitResp(lan string) *Resp {
	if lan == "" {
		lan = language.EnLan
	}
	RESP = &Resp{lan: lan}
	return RESP
}

// SubMsgType 子信息类型
type subMsgType map[int]string

const (
	// 系统
	Success           = 0
	UnknownError      = 10000
	ParameterError    = 10001
	OperationFailed   = 10002
	MissingParameters = 10003
	InvalidOperation  = 10004
	UserNotAuthorize  = 10005
	FrequencyTooFast  = 10006
	DataAlreadyExists = 10007
	DataDoesNotExist  = 10008
	UntrustedSource   = 10009

	// 用户
	UserExists    = 60000
	UserNotExists = 60001
	LackBalance   = 60002

	//验证
	PasswordError   = 70000
	TokenExpire     = 70002
	GoogleCodeError = 70005
)

// RespBody 响应体
type Elem struct {
	Ok   bool
	Code int
	Data interface{}
	Msg  string
}

// GetMsg 获取响应信息
func (r *Resp) GetMsg(b *Elem) (response gin.H) {
	if b.Ok && b.Code == 0 {
		b.Code = Success
	}
	msg := r.GetMsgStr(b.Code)
	if b.Msg != "" {
		msg += " | " + b.Msg
	}
	return gin.H{"code": b.Code, "message": msg, "data": b.Data}
}

func (r *Resp) SuccessRep(c *gin.Context, b *Elem) {
	c.JSON(http.StatusOK, r.GetMsg(b))
}

func (r *Resp) ParamErrRep(c *gin.Context, b *Elem) {
	c.JSON(http.StatusBadRequest, r.GetMsg(b))
}

func (r *Resp) ServerErrRep(c *gin.Context, b *Elem) {
	c.JSON(http.StatusInternalServerError, r.GetMsg(b))
}

func (r *Resp) GetMsgStr(Code int) (msg string) {
	var exist bool
	switch r.lan {
	case language.ZhCnLan:
		msg, exist = ZhCn[Code]
		if !exist {
			msg = ZhCn[UnknownError]
		}
	default:
		msg, exist = En[Code]
		if !exist {
			msg = En[UnknownError]
		}
	}
	return
}
