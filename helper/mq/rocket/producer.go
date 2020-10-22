package rocket

import (
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func (M *MQ) NewProducer(GroupName string) (p rocketmq.Producer, err error) {
	switch GroupName {
	case GroupToEmail:
	case GroupToTest:
	default:
		err = errors.New("unsupported type")
		return
	}
	p, err = rocketmq.NewProducer(
		// 集群
		producer.WithNameServer(M.ProducerAddr),
		// 单台
		//producer.WithNsResovler(primitive.NewPassthroughResolver(M.ProducerAddr)),
		producer.WithRetry(M.WithRetry),
		producer.WithGroupName(GroupName),
	)
	if err != nil {
		return
	}
	err = p.Start()
	return
}
