package price

import (
	"github.com/seperot/turtle-wear-api.git/getjson"
	"strconv"
)

type CoinValue struct {
	Usd string `json:"USD"`
	Btc string `json:"BTC"`
}

type CryptoFunc func(json getjson.RetrieveJson) string
type FiatFunc func(fiat string, json getjson.RetrieveJson) string

func Calc(cryptoFn CryptoFunc, fiatFn FiatFunc, json getjson.RetrieveJson) CoinValue {
	fiatPrice, _ := strconv.ParseFloat(fiatFn("USD", json), 64)
	exchangePrice, _ := strconv.ParseFloat(cryptoFn(json), 64)
	return CoinValue{"$" + strconv.FormatFloat(fiatPrice * exchangePrice, 'f', 9, 64),
		"₿" + strconv.FormatFloat(exchangePrice,'f',9, 64)}
}