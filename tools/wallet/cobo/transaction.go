package cobo

import (
	"strconv"
)

const (
	TransactionListUrl = "/v1/custody/transaction_history/"
	TransactionUrl     = "/v1/custody/transaction/"
)

type TransactionWhere struct {
	Coin      string
	Side      string
	Address   string
	MaxId     string
	MinId     string
	Limit     int
	BeginTime string
}

type Transaction struct {
	Id                  string
	Coin                string
	DisplayCode         string
	Description         string
	Decimal             int
	Address             string
	Memo                string
	SourceAddress       string
	SourceAddressDetail string
	Side                string
	Amount              string
	AbsAmount           string
	AbsCoboFee          string
	Txid                string
	VoutN               string
	RequestId           string
	Status              string
	CreatedTime         string
	LastTime            string
	ConfirmingThreshold int
	ConfirmedNum        int
	FeeCoin             string
	FeeAmount           int
	FeeDecimal          int
	Type                string
}

func (Co *Cobo) SetFindWhere(Condition *TransactionWhere) (params *Params) {
	var Par = Params{}
	params = &Par

	if Condition.Coin != "" {
		Par["coin"] = Condition.Coin
	}
	if Condition.Side != "" {
		Par["side"] = Condition.Side
	}
	if Condition.Address != "" {
		Par["address"] = Condition.Address
	}
	if Condition.MaxId != "" {
		Par["max_id"] = Condition.MaxId
	}
	if Condition.MinId != "" {
		Par["min_id"] = Condition.MinId
	}

	limit := strconv.Itoa(Condition.Limit)
	if limit != "" && Condition.Limit > 0 {
		Par["limit"] = limit
	}
	if Condition.BeginTime != "" {
		Par["begin_time"] = Condition.BeginTime
	}
	return
}

func (Co *Cobo) FindTransactionList(Condition *TransactionWhere) (Transaction []*Transaction, err error) {
	params := Co.SetFindWhere(Condition)
	//fmt.Println("params", params)
	result := Co.Request(GET, TransactionListUrl, *params)
	//fmt.Printf("result %+v \n", result)
	err = CheckBase(result, &Transaction)
	return
}

func (Co *Cobo) FindTransaction(CoboId string) (Transaction *Transaction, err error) {
	err = CheckBase(Co.Request(GET, TransactionUrl, Params{
		"id": CoboId,
	}), &Transaction)
	return
}
