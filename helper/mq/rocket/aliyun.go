package rocket

import (
	"fmt"
	"github.com/aliyunmq/mq-http-go-sdk"
	"github.com/gogap/errors"
	"strings"
	"time"
)

var AliyunMQClient mq_http_sdk.MQClient
var Prodrcer mq_http_sdk.MQProducer

func InitAliyun(endpoint, accessKey, secretKey, securityToken string) {
	tmpClient := mq_http_sdk.NewAliyunMQClient(endpoint, accessKey, secretKey, securityToken)
	AliyunMQClient = tmpClient
}

func NewAliyun() *AliyunMQ {
	return &AliyunMQ{
		Client:   AliyunMQClient,
		Producer: Prodrcer,
	}
}

type AliyunMQ struct {
	Client   mq_http_sdk.MQClient
	Producer mq_http_sdk.MQProducer
	Consumer mq_http_sdk.MQConsumer
}

func (M *AliyunMQ) InitProducer(instanceId, topic string) {
	M.Producer = AliyunMQClient.GetProducer(instanceId, topic)
}

func (M *AliyunMQ) PushMsg(Msg, Tag string) (resp mq_http_sdk.PublishMessageResponse, err error) {
	msg := mq_http_sdk.PublishMessageRequest{
		MessageBody: Msg, //消息内容
		MessageTag:  Tag, // 消息标签
	}
	return M.Producer.PublishMessage(msg)
}

func (M *AliyunMQ) InitConsumer(instanceId, topic, groupId, messageTag string) {
	M.Consumer = M.Client.GetConsumer(instanceId, topic, groupId, messageTag)
}

func (M *AliyunMQ) PullMsg(Business func([]mq_http_sdk.ConsumeMessageEntry) ([]string, error)) {
	endChan := make(chan int)
	respChan := make(chan mq_http_sdk.ConsumeMessageResponse)
	errChan := make(chan error)
	for {
		go func() {
			select {
			case resp := <-respChan:
				{
					handles, err := Business(resp.Messages)
					if err != nil {
						fmt.Printf("Business err %+v  -------->\n", err)
					} else {
						// NextConsumeTime前若不确认消息消费成功，则消息会重复消费
						// 消息句柄有时间戳，同一条消息每次消费拿到的都不一样
						ackerr := M.Consumer.AckMessage(handles)
						if ackerr != nil {
							// 某些消息的句柄可能超时了会导致确认不成功
							fmt.Println(ackerr)
							for _, errAckItem := range ackerr.(errors.ErrCode).Context()["Detail"].([]mq_http_sdk.ErrAckItem) {
								fmt.Printf("\tErrorHandle:%s, ErrorCode:%s, ErrorMsg:%s\n",
									errAckItem.ErrorHandle, errAckItem.ErrorCode, errAckItem.ErrorMsg)
							}
							time.Sleep(time.Duration(3) * time.Second)
						} else {
							fmt.Printf("Ack ---->\n\t%s\n", handles)
						}
					}
					endChan <- 1
				}
			case err := <-errChan:
				{
					// 没有消息
					if strings.Contains(err.(errors.ErrCode).Error(), "MessageNotExist") {
						fmt.Println("No new message, continue!")
					} else {
						fmt.Println(err)
						time.Sleep(time.Duration(3) * time.Second)
					}
					endChan <- 1
				}
			case <-time.After(35 * time.Second):
				{
					fmt.Println("Timeout of consumer message ??")
					endChan <- 1
				}
			}
		}()

		// 长轮询消费消息
		// 长轮询表示如果topic没有消息则请求会在服务端挂住3s，3s内如果有消息可以消费则立即返回
		M.Consumer.ConsumeMessage(respChan, errChan,
			3, // 一次最多消费3条(最多可设置为16条)
			3, // 长轮询时间3秒（最多可设置为30秒）
		)
		<-endChan
	}
}
