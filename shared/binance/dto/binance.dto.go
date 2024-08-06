package binance_dto

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
	TotalTradedQuotAssetVolumn string `json:"q"`
	StatisticsOpenTime         int64  `json:"O"`
	StatisticsClosTime         int64  `json:"C"`
	FirstTradeId               int64  `json:"F"`
	LastTradeId                int64  `json:"L"`
	TotalNumberOfTrades        int64  `json:"n"`
	AggregateTradeID           int64  `json:"a"`
	BestAskedQty               string `json:"A"`
	IsTheBuyerTheMarkerMaker   bool   `json:"m"`
	Ignore                     bool   `json:"M"`
	TradeTime                  int64  `json:"T"`
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
