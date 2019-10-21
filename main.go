package main

import (
	"fmt"
	"trtlFunctions/priceCalculator"
)

type coinValue struct {
	Usd string `json:"USD"`
	Btc string `json:"BTC"`
}

func show() (*coinValue, error) {
	bk := &coinValue{
		Usd: "$" + priceCalculator.PriceCalc("USD", priceCalculator.ExchangeOgre),
		Btc: "₿" + priceCalculator.PriceCalc("BTC", priceCalculator.ExchangeOgre),
	}
	return bk, nil
}

func main() {
	fmt.Println(show())
}