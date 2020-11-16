package Email

import (
	"fmt"
	"net/smtp"
	"strings"
)

var SMTP *Smtp

type Smtp struct {
	Host      string
	SendUser  string
	SendPwd   string
	ReplyUser string
	UserName  string
}

func New() *Smtp {
	return SMTP
}

func Init(Host, User, Pwd, Reply string) {
	SMTP = &Smtp{
		Host:      Host,
		SendUser:  User,
		SendPwd:   Pwd,
		ReplyUser: Reply,
	}
}

type SendData interface {
	// 接受方
	GetTo() string
	// 标题
	GetSubject() string
	// 内容
	GetBody() string
	// 内容类型
	GetMailType() string
}

func (E *Smtp) Send(Data SendData) error {
	hp := strings.Split(E.Host, ":")
	auth := smtp.PlainAuth("", E.SendUser, E.SendPwd, hp[0])
	var ContentType string
	if Data.GetMailType() == "html" {
		ContentType = "Content-Type: text/" + Data.GetMailType() + "; charset=UTF-8"
	} else {
		ContentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + Data.GetTo() + "\r\nFrom: " + E.ReplyUser + ">\r\nSubject: " + Data.GetSubject() + "\r\n" + ContentType + "\r\n\r\n" + Data.GetBody())
	sendTo := strings.Split(Data.GetTo(), ";")
	fmt.Println("Send Email ............")
	err := smtp.SendMail(E.Host, auth, E.SendUser, sendTo, msg)
	fmt.Println("Result ", err)
	return err
}
