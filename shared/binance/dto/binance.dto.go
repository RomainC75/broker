package binance_dto

import (
	"reflect"
	"shared/utils"
)

type BinanceAggTradeDto struct {
	EventType                  string `json:"e"`
	EventTime                  int64  `json:"E"`
	Symbol                     string `json:"s"`
	PriceChange                string `json:"p"`
	PriceChangePercent         string `json:"P"`
	OpenPrice                  string `json:"o"`
	HighPrice                  string `json:"h"`
	LastTradeID                int64  `json:"l"`
	LastPrice                  string `json:"c"`
	WeightedAveragePrice       string `json:"w"`
	TotalTradedBassAssetVolume string `json:"v"`
	TotalTradedQuotAssetVolume string `json:"q"`
	StatisticsOpenTime         int64  `json:"O"`
	StatisticsCloseTime        int64  `json:"C"`
	FirstTradeId               int64  `json:"F"`
	// LastTradeId                int64  `json:"L"`
	TotalNumberOfTrades      int64  `json:"n"`
	AggregateTradeID         int64  `json:"a"`
	BestAskedQty             string `json:"A"`
	IsTheBuyerTheMarkerMaker bool   `json:"m"`
	Ignore                   bool   `json:"M"`
	TradeTime                int64  `json:"T"`
}

type ReverseBinanceAggTradeDto struct {
	EventType                  string `json:"event_type"`
	EventTime                  int64  `json:"event_time"`
	Symbol                     string `json:"symbol"`
	PriceChange                string `json:"price_change"`
	PriceChangePercent         string `json:"price_change_percent,omitempty"`
	OpenPrice                  string `json:"open_price,omitempty"`
	HighPrice                  string `json:"high_price,omitempty"`
	LastTradeID                int64  `json:"last_trade_id"`
	LastPrice                  string `json:"last_price,omitempty"`
	WeightedAveragePrice       string `json:"weighted_average_price,omitempty"`
	TotalTradedBassAssetVolume string `json:"total_traded_bass_asset_volume,omitempty"`
	TotalTradedQuotAssetVolume string `json:"total_traded_quot_asset_volume"`
	StatisticsOpenTime         int64  `json:"statistics_open_time,omitempty"`
	StatisticsClosTime         int64  `json:"statistics_close_time,omitempty"`
	FirstTradeId               int64  `json:"first_trade_id"`
	TotalNumberOfTrades        int64  `json:"total_number_of_trades,omitempty"`
	AggregateTradeID           int64  `json:"aggregate_trade_id"`
	BestAskedQty               string `json:"bestAsked_qty,omitempty"`
	IsTheBuyerTheMarkerMaker   bool   `json:"is_the_buyer_the_marker_maker"`
	Ignore                     bool   `json:"ignore,omitempty"`
	TradeTime                  int64  `json:"trade_time"`
}

type BinanceDepthDto struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	TradeTime int64  `json:"T"`
	Symbol    string `json:"s"`
	Pair      string `json:"ps"`
	// Price Level , Qty
	Bids [2][]string `json:"b"`
	// PRice level to be , Qty
	AsksToBeUpdated [2][]string `json:"a"`
}

func ConvertToReverseBinanceAggTradeDto(bat BinanceAggTradeDto) (rBat ReverseBinanceAggTradeDto) {
	v1 := reflect.ValueOf(&bat).Elem()
	keys := getKeysOfStruct(v1)

	v2 := reflect.ValueOf(&rBat).Elem()

	for _, key := range keys {
		fieldValue := v1.FieldByName(key)
		targetValue := v2.FieldByName(key)
		if targetValue.IsValid() && targetValue.CanSet() {
			targetValue.Set(fieldValue)
		}
	}
	utils.PrettyDisplay("aggTrade ", rBat)
	return

}

func getKeysOfStruct(v reflect.Value) []string {
	keys := []string{}
	for i := 0; i < v.NumField(); i++ {
		keys = append(keys, v.Type().Field(i).Name)
	}
	return keys
}
