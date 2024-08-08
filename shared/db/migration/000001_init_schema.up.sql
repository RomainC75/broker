
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password  TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE currencies (
    id BIGSERIAL PRIMARY KEY,
    event_type TEXT NOT NULL,
    event_time DATE NOT NULL,
    -- insert enum ?? 
    symbol TEXT NOT NULL,
    price_change TEXT NOT NULL,
    last_trade_id BIGINT NOT NULL,
    total_traded_quot_asset_volume TEXT NOT NULL,
    first_trade_id BIGINT NOT NULL,
    aggregate_trade_id BIGINT NOT NULL,
    is_the_buyer_the_marker_maker BOOLEAN NOT NULL,
    ignor BOOLEAN NOT NULL,
    trade_time DATE NOT NULL
);

CREATE INDEX idx_symbol_eventTime ON currencies (symbol, event_time);