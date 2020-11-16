package RabbitMQModule

type BasePush interface {
	Verification() error
	GetMessage() (str string, err error)
}

type BasePull interface {
}
