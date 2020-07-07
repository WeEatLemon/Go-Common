package hoo

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	WithdrawUrl = "/api/open/vip/v1/withdraw"
)

type WithdrawWhere struct {
	CoinName        string
	TokenName       string
	OrderId         string
	Amount          string
	Fee             string
	ToAddress       string
	Memo            string
	ContractAddress string
	SendNum         int
}

func (H *Hoo) SetWithdrawWhere(Data *WithdrawWhere) (params *Params, err error) {
	if Data.TokenName == "" || Data.OrderId == "" || Data.Amount == "" || Data.ToAddress == "" {
		err = errors.New("参数有误！")
		return
	}

	var Par = Params{}
	params = &Par

	if Data.TokenName != "" {
		Par["token_name"] = Data.TokenName
	}

	if Data.OrderId != "" {
		Par["order_id"] = Data.OrderId
	}

	if Data.Amount != "" {
		Par["amount"] = Data.Amount
	}

	if Data.ToAddress != "" {
		Par["to_address"] = Data.ToAddress
	}

	if Data.ContractAddress != "" {
		Par["contract_address"] = Data.ContractAddress
	}
	return
}

type BaseResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    map[string]string
}

func (H *Hoo) SetWithdraw(Data *WithdrawWhere) (err error, hooResp string, hooData map[string]string) {
	params, err := H.SetWithdrawWhere(Data)
	if err != nil {
		return
	}
	fmt.Printf("Request Params is  %+v \n", params)
	hooResp, err = H.Request(POST, WithdrawUrl, *params)
	if err != nil {
		return
	}
	var data BaseResp
	_ = json.Unmarshal([]byte(hooResp), &data)
	hooData = data.Data
	if data.Code != "10000" {
		err = errors.New(data.Message)
	}
	return
}
