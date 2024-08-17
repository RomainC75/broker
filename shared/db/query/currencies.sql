-- -- name: GetStockWithIngredients :one
-- SELECT * FROM stocks
-- INNER JOIN ingredients ON stocks.id = ingredients.id
-- WHERE stocks.user_id = $1 AND stocks.name = $2 LIMIT 1;

-- name: GetLastXCurrencies :many
SELECT * FROM currencies
WHERE symbol = $1
ORDER BY event_time DESC
LIMIT $2;

-- name: CreateStock :one
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
RETURNING *;

-- name: UpdateStock :one
-- UPDATE stocks
-- SET 
--     name = $3,
--     updated_at = $4
-- WHERE id = $1 AND user_id = $2
-- RETURNING *;

-- name: DeleteStock :many
-- DELETE FROM stocks
-- WHERE id = $1 AND user_id = $2 
-- RETURNING *;