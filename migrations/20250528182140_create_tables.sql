-- +goose Up
-- +goose StatementBegin
CREATE TABLE transaction
(
    id          SERIAL PRIMARY KEY,
    amount      NUMERIC(12, 2) NOT NULL,
    name        TEXT           NOT NULL,
    note        TEXT,
    category_id BIGINT         NOT NULL,
    user_id     BIGINT         NOT NULL,
    created_at  TIMESTAMPTZ    NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ    NOT NULL DEFAULT now()
);

CREATE TABLE recurring_transaction
(
    id                  SERIAL PRIMARY KEY,
    transaction_id      INTEGER     NOT NULL REFERENCES transaction (id) ON DELETE CASCADE,
    recurrence_interval INTERVAL    NOT NULL,
    is_active           BOOLEAN     NOT NULL DEFAULT TRUE,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE period_financial_summary
(
    id                   SERIAL PRIMARY KEY,
    date_from            DATE           NOT NULL,
    date_to              DATE           NOT NULL,
    starting_balance     NUMERIC(12, 2) NOT NULL,
    projected_balance    NUMERIC(12, 2),
    actual_balance       NUMERIC(12, 2),
    projected_net_change NUMERIC(12, 2),
    actual_net_change    NUMERIC(12, 2),
    created_at           TIMESTAMPTZ    NOT NULL DEFAULT now(),
    updated_at           TIMESTAMPTZ    NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE period_financial_summary;
DROP TABLE recurring_transaction;
DROP TABLE transaction;
-- +goose StatementEnd
