package Email

import (
	"fmt"
	EmailModules "github.com/WeEatLemon/Go-Common/helper/email/modules"
	"github.com/WeEatLemon/Go-Common/language"
	"testing"
	"time"
)

func InitDb() {
	Host := "smtp.163.com:465"
	User := "xxx@163.com"
	Pwd := ""
	Reply := "xxx@163.com"

	Init(Host, User, Pwd, Reply)
}

func TestSmtp_Send(t *testing.T) {
	InitDb()
	Data := &EmailModules.Registered{
		User:       "269119257@qq.com",
		Platform:   "sofa",
		Expiration: time.Now(),
		Language:   language.ZhCnLan,
	}
	err := SMTP.Send(Data)
	fmt.Println("err", err)
}
