package hoo

import "strconv"

const (
	AccountsUrl   = "/api/open/vip/v1/accounts"
	NewAddressUrl = "/api/open/vip/v1/address"
)

type AccountWhere struct {
	CoinName string
}

func (H *Hoo) AccountWhere(Condition *AccountWhere) (params *Params) {
	var Par = Params{}
	params = &Par
	if Condition.CoinName != "" {
		Par["coin_name"] = Condition.CoinName
	}
	return
}

type Account struct {
	Balance   string `json:"balance"`
	Frozen    string `json:"frozen"`
	TokenName string `json:"token_name"`
}

func (H *Hoo) GetAccount(Data *AccountWhere) (Account *Account, err error) {
	params := H.AccountWhere(Data)
	data, err := H.Request(POST, AccountsUrl, *params)
	if err != nil {
		return
	}
	err = CheckBase(data, &Account)
	return
}

type AddressWhere struct {
	CoinName string
	Num      int
}

func (H *Hoo) AddressWhere(Condition *AddressWhere) (params *Params) {
	var Par = Params{}
	params = &Par
	if Condition.CoinName != "" {
		Par["coin_name"] = Condition.CoinName
	}
	Num := strconv.Itoa(Condition.Num)
	if Num != "" && Condition.Num > 0 {
		Par["num"] = Num
	}
	return
}
func (H *Hoo) NewAddress(Data *AddressWhere) (Address []string, err error) {
	params := H.AddressWhere(Data)
	data, err := H.Request(POST, NewAddressUrl, *params)
	if err != nil {
		return
	}
	err = CheckBase(data, &Address)
	return
}
