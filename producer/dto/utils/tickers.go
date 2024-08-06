package binance_utils

import (
	"fmt"
	"strings"
)

//* Example : ethusdt@aggTrade

func GetAggTradeTickers(tickersChain string) []string {
	tickers := strings.Split(tickersChain, ",")
	for i, ticker := range tickers {
		tickers[i] = fmt.Sprintf("%s@aggTrade", ticker)
	}
	return tickers
}
