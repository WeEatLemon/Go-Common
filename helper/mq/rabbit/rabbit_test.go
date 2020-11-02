package RabbitMQ

import (
	"encoding/json"
	"fmt"
	"github.com/WeEatLemon/Go-Common/helper"
	"github.com/streadway/amqp"
	"testing"
	"time"
)

func InitRab() {
	Rab := InitRabbitMQ("root", "root", "127.0.0.1:5672")

	err := Rab.NewQueue("Notice", "Notice", "ST-Notice")

	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("new MQ is success")
	}
}

type TestMq struct {
	Type       string `json:"type"`
	Addr       string `json:"address"`
	NoticeType string `json:"notice_type"`
}

func TestRabbitMQ_ConsumeSimple(t *testing.T) {
	InitRab()
	Rab := NewRabbitMQ()
	go func() {
		for {
			PublishSimple(Rab)
			time.Sleep(time.Second * 3)
		}
	}()
	Rab.ConsumeSimple(Business)
}

func PublishSimple(Rab *RabbitMQ) {
	Msg := TestMq{
		Type:       "phone",
		Addr:       helper.GetUuid(),
		NoticeType: "Registered",
	}
	m, _ := json.Marshal(Msg)
	err := Rab.PublishSimple(string(m))
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("PublishSimple is success")
	}
}

func Business(Message <-chan amqp.Delivery) {
	for Msg := range Message {
		var data TestMq
		err := json.Unmarshal(Msg.Body, &data)
		if err != nil {
			fmt.Println("err is", err)
		} else {
			fmt.Printf("pull msg is %+v \n", data)
		}
	}
}
