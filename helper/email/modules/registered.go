package EmailModules

import (
	"github.com/WeEatLemon/Go-Common/language"
	"time"
)

type Registered struct {
	User       string    `json:"user"`
	Platform   string    `json:"platform"`
	Code       string    `json:"code"`
	Expiration time.Time `json:"expiration"`
	Language   string    `json:"language"`
}

func (M *Registered) GetTo() string {
	return M.User
}

func (M *Registered) GetSubject() string {
	var subject string
	switch M.Language {
	case language.EnLan:
		subject = "Registered for" + M.Platform
	default:
		subject = M.Platform + "欢迎您"
	}
	return subject
}

func (M *Registered) GetBody() string {
	var body string
	switch M.Language {
	case language.EnLan:
		body += "Your verification code is " + M.Code
		body += "Valid until " + M.Expiration.Format("2006-01-02 15:04:05")
	default:
		body += "您的验证码为 " + M.Code
		body += "有效期至 " + M.Expiration.Format("2006-01-02 15:04:05")
	}
	return body
}

func (M *Registered) GetMailType() string {
	return "text"
}
