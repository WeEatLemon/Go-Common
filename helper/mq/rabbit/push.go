package RabbitMQ

import "github.com/IEatLemons/GoHelper/helper/mq/rabbit/modules"

func PushMsg(R *RabbitMQ, Data RabbitMQModule.BasePush) (err error) {
	err = Data.Verification()
	if err != nil {
		return err
	}
	Msg, err := Data.GetMessage()
	if err != nil {
		return err
	}
	err = R.PublishSimple(Msg)
	return
}
