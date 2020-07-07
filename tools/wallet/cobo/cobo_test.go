package cobo

import (
	"fmt"
	"testing"
)

func TestGenerateRandomKeyPair(t *testing.T) {
	key := GenerateRandomKeyPair()
	fmt.Printf("%+v \n", key)
}

const (
	ApiKey    = "0321241c80f605cd6265e198f102b9943113af60ccbe42572ee0ad34df9417f413"
	ApiSecret = "7106c3885cddbf3a85397034f621d6a89e2f2041edc18319f3c0d36ec3cfb35c"
	HOST      = "https://api.sandbox.cobo.com"
	CoboPub   = "032f45930f652d72e0c90f71869dfe9af7d713b1f67dc2f7cb51f9572778b9c876"
)

var Co = NewCobo(ApiKey, ApiSecret, HOST, CoboPub)

func TestRequest(t *testing.T) {

	result := Co.Request(GET, OrgInfoUrl, Params{})
	fmt.Println("result", result)
}

func TestCobo_GetGetAccountList(t *testing.T) {
	list, err := Co.GetAccountList()

	fmt.Printf("err : %+v \n", err)

	fmt.Printf("list : %+v \n", list)
}

func TestCobo_GetAccountInfo(t *testing.T) {
	Coin := "XTN"

	info, err := Co.GetAccountInfo(Coin)

	fmt.Printf("err : %+v \n", err)

	fmt.Printf("info : %+v \n", info)
}

func TestCobo_NewAddress(t *testing.T) {
	Coin := "XTN"

	Address, err := Co.NewAddress(Coin)

	fmt.Printf("err : %+v \n", err)

	fmt.Printf("info : %+v \n", Address)
}

func TestCobo_AddressHistory(t *testing.T) {
	Coin := "XTN"

	list, err := Co.AddressHistory(Coin)

	fmt.Printf("err : %+v \n", err)

	for k, v := range list {
		fmt.Printf("k : %+v \n", k)
		fmt.Printf("v : %+v \n", v)
	}
}

func TestCobo_ValidAddress(t *testing.T) {
	Coin := "XTN"
	Address := "0x5a65eeae67c42140602bf1fae46a0768a4e592c6"

	result, err := Co.ValidAddress(Coin, Address)

	fmt.Printf("err : %+v \n", err)

	fmt.Printf("result : %+v \n", result)
}

func TestCobo_FindTransaction(t *testing.T) {

	list, err := Co.FindTransactionList(&TransactionWhere{})

	fmt.Printf("err : %+v \n", err)

	fmt.Printf("list : %+v \n", list)

	for _, v := range list {
		tr, err := Co.FindTransaction(v.Id)
		fmt.Printf("err : %+v \n", err)
		fmt.Printf("tr : %+v \n", tr)
	}
}
