package cobo

const (
	OrgInfoUrl        = "/v1/custody/org_info/"
	CoinInfoUrl       = "/v1/custody/coin_info/"
	NewAddressUrl     = "/v1/custody/new_address/"
	AddressHistoryUrl = "/v1/custody/address_history/"
	ValidAddress      = "/v1/custody/is_valid_address/"
)

type OrgInfo struct {
	Name   string     `json:"name" binding:"name"`
	Assets []CoinInfo `json:"assets" binding:"assets"`
}

func (Co *Cobo) GetAccountList() (OrgInfoList *OrgInfo, err error) {
	err = CheckBase(Co.Request(GET, OrgInfoUrl, Params{}), &OrgInfoList)
	return
}

type CoinInfo struct {
	Coin                string `json:"coin" binding:"coin"`
	DisplayCode         string `json:"display_code" binding:"display_code"`
	Description         string `json:"description" binding:"description"`
	Decimal             int    `json:"decimal" binding:"decimal"`
	CanDeposit          bool   `json:"can_deposit" binding:"can_deposit"`
	CanWithdraw         bool   `json:"can_withdraw" binding:"can_withdraw"`
	Balance             string `json:"balance" binding:"balance"`
	AbsBalance          string `json:"abs_balance" binding:"abs_balance"`
	FeeCoin             string `json:"fee_coin" binding:"fee_coin"`
	AbsEstimateFee      string `json:"abs_estimate_fee" binding:"abs_estimate_fee"`
	ConfirmingThreshold int    `json:"confirming_threshold" binding:"confirming_threshold"`
	DustThreshold       int    `json:"dust_threshold" binding:"dust_threshold"`
	TokenAddress        string `json:"token_address" binding:"token_address"`
	RequireMemo         string `json:"require_memo" binding:"require_memo"`
}

func (Co *Cobo) GetAccountInfo(Coin string) (CoinInfo *CoinInfo, err error) {
	result := Co.Request(GET, CoinInfoUrl, Params{
		"coin": Coin,
	})
	err = CheckBase(result, &CoinInfo)
	return
}

type Address struct {
	Coin    string `json:"coin" binding:"coin"`
	Address string `json:"address" binding:"address"`
}

func (Co *Cobo) NewAddress(Coin string) (Address *Address, err error) {
	err = CheckBase(Co.Request(POST, NewAddressUrl, Params{
		"coin": Coin,
	}), &Address)
	return
}

func (Co *Cobo) AddressHistory(Coin string) (AddressList []*Address, err error) {
	err = CheckBase(Co.Request(GET, AddressHistoryUrl, Params{
		"coin": Coin,
	}), &AddressList)
	return
}

func (Co *Cobo) ValidAddress(Coin, Address string) (result bool, err error) {
	err = CheckBase(Co.Request(GET, ValidAddress, Params{
		"coin":    Coin,
		"address": Address,
	}), &result)
	return
}
