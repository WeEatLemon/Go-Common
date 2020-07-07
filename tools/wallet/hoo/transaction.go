package hoo

import "fmt"

const (
	KlineUrl = "/open/v1/kline"
)

type KlineWhere struct {
	Symbol string
	Type   string
}

type Kline struct {
	Amount string
	Close  string
	High   string
	Low    string
	Open   string
	Time   string
	Volume string
}

func (H *Hoo) GetKline(Data *KlineWhere) (result *Kline, err error) {
	params, err := H.SetKlineWhere(Data)
	if err != nil {
		return
	}
	fmt.Printf("Request Params is  %+v \n", params)
	reData, err := H.Request(GET, KlineUrl, *params)
	if err != nil {
		return
	}
	fmt.Println("CheckBase....", reData)
	err = CheckBase(reData, &result)
	return
}

func (H *Hoo) SetKlineWhere(Data *KlineWhere) (params *Params, err error) {
	var Par = Params{}
	params = &Par
	if Data.Symbol != "" {
		Par["symbol"] = Data.Symbol
	}

	if Data.Type != "" {
		Par["type"] = Data.Type
	}
	return
}
