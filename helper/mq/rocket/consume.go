package rocket

import (
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
)

func (M *MQ) NewPushConsume(GroupName string) (c rocketmq.PushConsumer, err error) {
	switch GroupName {
	case GroupToEmail:
	case GroupToTest:
	default:
		err = errors.New("unsupported type")
		return
	}
	c, err = rocketmq.NewPushConsumer(
		consumer.WithNameServer(M.ProducerAddr),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(GroupName),
	)
	return
}
