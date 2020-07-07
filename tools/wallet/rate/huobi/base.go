package huobi

import (
	"fmt"
	"github.com/IEatLemons/GoBase/common/tools/wallet/rate"
)

const (
	Api = "https://api.binance.com/api/v3/ticker/price"

	BtcToUsdt = "BTCUSDT"
)

type Params map[string]string

func GetRate() {
	params := Params{}
	params["symbol"] = BtcToUsdt

	data, err := rate.Request(rate.GET, Api, params)

	fmt.Printf("err is %+v \n", err)

	fmt.Printf("result is %+v \n", data)
}
