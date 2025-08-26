-- +goose Up
-- +goose StatementBegin
CREATE TABLE finance_detailing
(
    id                 SERIAL PRIMARY KEY,
    user_id            BIGINT         NOT NULL,
    date_from          TIMESTAMPTZ    NOT NULL,
    date_to            TIMESTAMPTZ    NOT NULL,
    initial_amount     NUMERIC(12, 2) NOT NULL,
    current_amount     NUMERIC(12, 2) NOT NULL,
    total_income       NUMERIC(12, 2) NOT NULL,
    total_expense      NUMERIC(12, 2) NOT NULL,
    profit_estimated   NUMERIC(12, 2) NOT NULL,
    profit_real        NUMERIC(12, 2) NOT NULL,
    after_amount_net   NUMERIC(12, 2) NOT NULL,
    after_amount_gross NUMERIC(12, 2) NOT NULL,
    created_at         TIMESTAMPTZ    NOT NULL DEFAULT now(),
    updated_at         TIMESTAMPTZ    NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE finance_detailing;
-- +goose StatementEnd