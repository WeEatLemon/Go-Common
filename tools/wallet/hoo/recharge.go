package hoo

import "fmt"

const (
	OrdersUrl = "/api/open/vip/v1/orders"
)

type OrderWhere struct {
	CoinName  string
	TradeType int
	StartAt   int
	EndAt     int
	Pagenum   int
	Pagesize  int
}

type Orders struct {
	Count   string
	Pagenum string
	Records []Order
}

type Order struct {
	OuterOrderNo  string
	TradeType     string
	CoinName      string
	ChainName     string
	TransactionId string
	BlockHeight   string
	Confirmations string
	FromAddress   string
	TtoAddress    string
	Amount        string
	Fee           string
	Status        string
	CreateAt      string
	ProcessAt     string
}

func (H *Hoo) GetOrders(Data *OrderWhere) (result bool, err error) {
	params, err := H.SetOrdersWhere(Data)
	if err != nil {
		return
	}
	fmt.Printf("Request Params is  %+v \n", params)
	reData, err := H.Request(POST, WithdrawUrl, *params)
	if err != nil {
		return
	}
	fmt.Println("CheckBase....", reData)
	err = CheckBase(reData, &result)
	return
}

func (H *Hoo) SetOrdersWhere(Data *OrderWhere) (params *Params, err error) {
	var Par = Params{}
	params = &Par
	if Data.CoinName != "" {
		Par["coin_name"] = Data.CoinName
	}
	return
}
