package config_utils

import (
	"fmt"
	"strings"
)

//* Example : ethusdt@aggTrade

func GetAggTradesFromTickers(tickers []string) []string {
	aggTrades := []string{}
	for _, ticker := range tickers {
		aggTrades = append(aggTrades, fmt.Sprintf("%s@aggTrade", ticker))
	}
	return aggTrades
}

func SeparateTickers(tickersChain string) []string {
	return strings.Split(tickersChain, ",")
}
