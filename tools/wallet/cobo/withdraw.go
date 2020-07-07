package cobo

import "errors"

const (
	SendWithdrawUrl = "/v1/custody/new_withdraw_request/"
)

type RequestData struct {
	Coin          string
	RequestId     string
	Address       string
	Amount        string
	Memo          string
	ForceExternal string
	ForceInternal string
}

func (Co *Cobo) SetRequestData(Data *RequestData) (params *Params, err error) {
	if Data.Coin == "" && Data.RequestId == "" && Data.Address == "" && Data.Amount == "" {
		err = errors.New("缺少必要参数")
		return
	}

	Par := Params{}
	params = &Par

	Par["coin"] = Data.Coin
	Par["request_id"] = Data.RequestId
	Par["address"] = Data.Address
	Par["amount"] = Data.Amount

	if Data.Memo != "" {
		Par["memo"] = Data.Memo
	}

	return
}

func (Co *Cobo) SendWithdraw(Data *RequestData) (result bool, err error) {
	params, err := Co.SetRequestData(Data)
	if err != nil {
		return
	}
	err = CheckBase(Co.Request(POST, SendWithdrawUrl, *params), &result)
	return
}
