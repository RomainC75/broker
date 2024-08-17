package db

import (
	"context"
	"time"
)

const createStock = `-- name: CreateStock :one
INSERT INTO currencies (
    event_type,
    event_time,
    symbol,
    price_change,
    last_trade_id,
    total_traded_quot_asset_volume,
    first_trade_id,
    aggregate_trade_id,
    is_the_buyer_the_marker_maker,
    is_ignore,
    trade_time
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING id, event_type, event_time, symbol, price_change, last_trade_id, total_traded_quot_asset_volume, first_trade_id, aggregate_trade_id, is_the_buyer_the_marker_maker, is_ignore, trade_time
`

type CreateStockParams struct {
	EventType                  string    `json:"eventType"`
	EventTime                  time.Time `json:"eventTime"`
	Symbol                     string    `json:"symbol"`
	PriceChange                string    `json:"priceChange"`
	LastTradeID                int64     `json:"lastTradeId"`
	TotalTradedQuotAssetVolume string    `json:"totalTradedQuotAssetVolume"`
	FirstTradeID               int64     `json:"firstTradeId"`
	AggregateTradeID           int64     `json:"aggregateTradeId"`
	IsTheBuyerTheMarkerMaker   bool      `json:"isTheBuyerTheMarkerMaker"`
	IsIgnore                   bool      `json:"isIgnore"`
	TradeTime                  time.Time `json:"tradeTime"`
}

func (q *Queries) CreateStock(ctx context.Context, arg CreateStockParams) (Currency, error) {
	row := q.db.QueryRowContext(ctx, createStock,
		arg.EventType,
		arg.EventTime,
		arg.Symbol,
		arg.PriceChange,
		arg.LastTradeID,
		arg.TotalTradedQuotAssetVolume,
		arg.FirstTradeID,
		arg.AggregateTradeID,
		arg.IsTheBuyerTheMarkerMaker,
		arg.IsIgnore,
		arg.TradeTime,
	)
	var i Currency
	err := row.Scan(
		&i.ID,
		&i.EventType,
		&i.EventTime,
		&i.Symbol,
		&i.PriceChange,
		&i.LastTradeID,
		&i.TotalTradedQuotAssetVolume,
		&i.FirstTradeID,
		&i.AggregateTradeID,
		&i.IsTheBuyerTheMarkerMaker,
		&i.IsIgnore,
		&i.TradeTime,
	)
	return i, err
}

const getLastXCurrencies = `-- name: GetLastXCurrencies :many

SELECT id, event_type, event_time, symbol, price_change, last_trade_id, total_traded_quot_asset_volume, first_trade_id, aggregate_trade_id, is_the_buyer_the_marker_maker, is_ignore, trade_time FROM currencies
WHERE symbol = $1
ORDER BY event_time DESC
LIMIT $2
`

type GetLastXCurrenciesParams struct {
	Symbol string `json:"symbol"`
	Limit  int32  `json:"limit"`
}

// -- name: GetStockWithIngredients :one
// SELECT * FROM stocks
// INNER JOIN ingredients ON stocks.id = ingredients.id
// WHERE stocks.user_id = $1 AND stocks.name = $2 LIMIT 1;
func (q *Queries) GetLastXCurrencies(ctx context.Context, arg GetLastXCurrenciesParams) ([]Currency, error) {
	rows, err := q.db.QueryContext(ctx, getLastXCurrencies, arg.Symbol, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Currency{}
	for rows.Next() {
		var i Currency
		if err := rows.Scan(
			&i.ID,
			&i.EventType,
			&i.EventTime,
			&i.Symbol,
			&i.PriceChange,
			&i.LastTradeID,
			&i.TotalTradedQuotAssetVolume,
			&i.FirstTradeID,
			&i.AggregateTradeID,
			&i.IsTheBuyerTheMarkerMaker,
			&i.IsIgnore,
			&i.TradeTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
