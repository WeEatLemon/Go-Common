package wechat

import (
	"fmt"
	"testing"
)

func TestRMPay(t *testing.T) {
	params := &ThirdPayToRM{
		Platform:    "pin-min",
		OrderSn:     "test4",
		Token:       "TZcJ0hB9LnMOS3r9hrPqP9o9Upmhehhv",
		PaymentType: 11,
		PayAmount:   0.01,
		OrderType:   1,
		Lang:        "en",
		Client:      1,
		ProcessType: 4,
	}

	result, err := NewRm().GetPayUrl(params)
	fmt.Printf("err is %+v \n", err)
	fmt.Printf("result is %+v \n", result)
}
