-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders
(
  id         SERIAL      PRIMARY KEY NOT NULL,
  user_id    BIGINT      NOT NULL,
  status     TEXT        NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS items_orders
(
  sku        BIGINT      NOT NULL,
  order_id   BIGINT      NOT NULL,
  item_count BIGINT      NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (sku, order_id)
);

CREATE TABLE IF NOT EXISTS stocks
(
  sku        BIGINT      PRIMARY KEY NOT NULL,
  total      BIGINT      NOT NULL DEFAULT 0,
  reserved   BIGINT      NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS items_orders;
DROP TABLE IF EXISTS stocks;
-- +goose StatementEnd
