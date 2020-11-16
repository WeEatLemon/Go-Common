package RabbitMQ

import (
	"encoding/json"
	"fmt"
	"github.com/WeEatLemon/Go-Common/helper"
	"github.com/WeEatLemon/Go-Common/helper/email"
	"github.com/WeEatLemon/Go-Common/helper/email/modules"
	"github.com/WeEatLemon/Go-Common/helper/mq/rabbit/modules"
	"github.com/WeEatLemon/Go-Common/middle"
	"github.com/streadway/amqp"
	"testing"
	"time"
)

var ES *Email.Smtp

func InitRab() {
	Rab := InitRabbitMQ("test", "test", "120.24.25.253:5672")
	//defer Rab.Destroy()
	err := Rab.NewQueue("Notice", "Notice", "ST-Notice")

	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("new MQ is success")

		Host := "smtp.163.com:465"
		User := "xxx@163.com"
		Pwd := ""
		Reply := "xxx@163.com"
		Email.Init(Host, User, Pwd, Reply)
	}
}

func TestRabbitMQ_ConsumeSimple(t *testing.T) {
	InitRab()
	Rab := NewRabbitMQ()
	go func() {
		for i := 0; i < 10; i++ {
			PublishSimple(Rab)
		}
	}()
	Rab.ConsumeSimple(Business)
}

func PublishSimple(Rab *RabbitMQ) {
	Msg := &RabbitMQModule.Registered{
		User:       "269119257@qq.com",
		Code:       helper.GetUuid(),
		Platform:   middle.Default,
		Expiration: time.Now().Add(time.Minute * 5),
	}
	err := PushMsg(Rab, Msg)
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("PublishSimple is success")
	}
}

func Business(Message <-chan amqp.Delivery) {
	ES = Email.New()
	for Msg := range Message {
		var data EmailModules.Registered
		err := json.Unmarshal(Msg.Body, &data)
		if err != nil {
			fmt.Println("Unmarshal err is", err)
		} else {
			fmt.Printf("data is %+v \n", data)
			err = ES.Send(&data)
			if err != nil {
				fmt.Println("Send err is", err)
			} else {
				fmt.Println("Send success")
			}
		}
	}
}
