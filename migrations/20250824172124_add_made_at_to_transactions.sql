-- +goose Up
-- +goose StatementBegin
ALTER TABLE transactions
    ADD COLUMN made_at TIMESTAMPTZ NOT NULL DEFAULT now();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE transactions
    DROP COLUMN made_at;
-- +goose StatementEnd