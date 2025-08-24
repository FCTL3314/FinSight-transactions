-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions
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

CREATE TABLE recurring_transactions
(
    id                  SERIAL PRIMARY KEY,
    transaction_id      INTEGER     NOT NULL REFERENCES transactions (id) ON DELETE CASCADE,
    recurrence_interval INTERVAL    NOT NULL,
    is_active           BOOLEAN     NOT NULL DEFAULT TRUE,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE recurring_transactions;
DROP TABLE transactions;
-- +goose StatementEnd