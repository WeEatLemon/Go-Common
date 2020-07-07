package hoo

import (
	"fmt"
	"testing"
)

const (
	ClientId     = ""
	ClientSecret = ""
	Host         = ""
)

var H = NewHoo(ClientId, ClientSecret, Host)

func TestGetAccount(t *testing.T) {
	p := &AccountWhere{
		CoinName: "BTC",
	}
	fmt.Printf("params is %+v \n", p)
	Account, err := H.GetAccount(p)
	fmt.Println("err", err)
	fmt.Println("Account TokenName ", Account.TokenName)
	fmt.Println("Account Balance ", Account.Balance)
	fmt.Println("Account Frozen ", Account.Frozen)
}

func TestNewAddress(t *testing.T) {
	p := &AddressWhere{
		CoinName: "ETH",
		Num:      1,
	}
	fmt.Printf("params is %+v \n", p)
	Address, err := H.NewAddress(p)
	fmt.Println("err", err)
	fmt.Println("Address ", Address)
}

func TestWithdraw(t *testing.T) {
	p := &WithdrawWhere{
		CoinName:        "USDT",
		TokenName:       "USDT-ERC20",
		OrderId:         "sup-min26934277t",
		Amount:          "0.01",
		ToAddress:       "0x00f49805af7c2e9c18f024b4cb3c94a49a1ff2d1",
		ContractAddress: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
	}
	err, hooResp, hooData := H.SetWithdraw(p)
	fmt.Println("err", err)
	fmt.Println("Address ", hooResp)
	fmt.Println("Address ", hooData)
}
