package db

import (
	"time"
)

type Currency struct {
	ID                         int64     `json:"id"`
	EventType                  string    `json:"eventType"`
	EventTime                  time.Time `json:"eventTime"`
	Symbol                     string    `json:"symbol"`
	PriceChange                string    `json:"priceChange"`
	LastTradeID                int64     `json:"lastTradeId"`
	TotalTradedQuotAssetVolume string    `json:"totalTradedQuotAssetVolume"`
	FirstTradeID               int64     `json:"firstTradeId"`
	AggregateTradeID           int64     `json:"aggregateTradeId"`
	IsTheBuyerTheMarkerMaker   bool      `json:"isTheBuyerTheMarkerMaker"`
	Ignor                      bool      `json:"ignor"`
	TradeTime                  time.Time `json:"tradeTime"`
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
