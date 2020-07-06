package request

import (
	"fmt"
	"testing"
)

func TestRequest(t *testing.T) {
	//params := make(ReqParams)
	var params = ReqParams{}
	params["order_no"] = "20200703104044Kx5DHS"
	params["pay_currency"] = "CNY"
	params["pay_protocol"] = ""
	params["pay_type"] = "WeChat"
	params["pay_status"] = "success"
	params["pay_amount"] = "0.01"

	ret, err := Request("PUT", "https://pinapi.rockminer.com/pin/order/notice", params, "")
	fmt.Println(ret, err)
}
