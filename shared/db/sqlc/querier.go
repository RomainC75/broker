package db

import (
	"context"
)

type Querier interface {
	CreateStock(ctx context.Context, arg CreateStockParams) (Currency, error)
	// -- name: GetStockWithIngredients :one
	// SELECT * FROM stocks
	// INNER JOIN ingredients ON stocks.id = ingredients.id
	// WHERE stocks.user_id = $1 AND stocks.name = $2 LIMIT 1;
	GetLastXCurrencies(ctx context.Context, arg GetLastXCurrenciesParams) ([]Currency, error)
}

var _ Querier = (*Queries)(nil)
