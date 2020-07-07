package hoo

import "strings"

func FilterTransactionData(params *Transaction) {
	CoinName := strings.Split(params.CoinName, "-")
	params.Currency = CoinName[0]
	if len(CoinName) > 1 {
		params.Protocol = CoinName[1]
	}
	switch params.Message {
	case "confirming":
		params.Status = "confirm"
	case "success":
		params.Status = "success"
	default:
		params.Status = "failure"
	}
}
