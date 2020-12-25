package RabbitMQ

import (
	"errors"
	"fmt"
	"github.com/IEatLemons/GoHelper/helper"
	"github.com/streadway/amqp"
	"log"
	"time"
)

var R *RabbitMQ

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// Key
	Key string
	// 连接信息
	Link string
}

func InitRabbitMQ(user, pwd, host string) *RabbitMQ {
	R = &RabbitMQ{
		Link: "amqp://" + user + ":" + pwd + "@" + host,
	}
	return R
}

func NewRabbitMQ() *RabbitMQ {
	return R
}

// NewRabbitMQ 创建结构体实例
func (r *RabbitMQ) NewQueue(queueName, exchange, key string) (err error) {
	if r.Link == "" {
		err = errors.New("please init RabbitMQ")
		return
	}
	r.QueueName = queueName
	r.Exchange = exchange
	r.Key = key
	// 创建rabbitmq连接
	r.conn, err = amqp.Dial(r.Link)
	if err != nil {
		return
	}

	r.channel, err = R.conn.Channel()
	return
}

func (r *RabbitMQ) Destroy() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

func (r *RabbitMQ) QueueDeclare() (amqp.Queue, error) {
	// 1. 申请队列，如果队列不存在会自动创建，如何存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	return r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		true,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
}

func (r *RabbitMQ) PublishSimple(message string) (err error) {
	_, err = r.QueueDeclare()
	if err != nil {
		return
	}

	// 2.发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		// 如果为true, 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		true,
		// 如果为true, 当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			MessageId:   helper.GetUuid(),
			Timestamp:   time.Time{},
			Body:        []byte(message),
		},
	)
	return
}

// ConsumeSimple 使用 goroutine 消费消息
func (r *RabbitMQ) ConsumeSimple(Business func(<-chan amqp.Delivery)) {
	_, err := r.QueueDeclare()
	if err != nil {
		fmt.Println(err)
	}

	// 接收消息
	messages, err := r.channel.Consume(
		r.QueueName,
		// 用来区分多个消费者
		"",
		// 是否自动应答
		true,
		// 是否具有排他性
		false,
		// 如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		// 队列消费是否阻塞
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 启用协和处理消息
	go func() {
		Business(messages)
	}()

	log.Printf("[*] Waiting for message, To exit press CTRL+C")
	<-forever
}
