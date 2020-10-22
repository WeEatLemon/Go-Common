package rocket

import (
	"encoding/json"
	"errors"
)

var Rocket *MQ

const (
	GroupToTest  = "Test:"
	GroupToEmail = "Email:"
)

type MQ struct {
	ProducerAddr []string
	WithRetry    int
}

func Init(ProducerAddress string, WithRetry int) (err error) {
	var ProducerAddr []string
	err = json.Unmarshal([]byte(ProducerAddress), &ProducerAddr)
	if err != nil {
		err = errors.New("")
		return
	}
	Rocket = &MQ{
		ProducerAddr: ProducerAddr,
		WithRetry:    WithRetry,
	}
	return
}

func New() *MQ {
	return Rocket
}
