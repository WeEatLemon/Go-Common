package RabbitMQModule

import (
	"encoding/json"
	"time"
)

type Registered struct {
	User       string    `json:"user"`
	Platform   string    `json:"platform"`
	Code       string    `json:"code"`
	Expiration time.Time `json:"expiration"`
	Language   string    `json:"language"`
}

func (M *Registered) GetMessage() (str string, err error) {
	tmp, err := json.Marshal(M)
	if err != nil {
		return
	}
	str = string(tmp)
	return
}

func (M *Registered) Verification() (err error) {

	return
}
