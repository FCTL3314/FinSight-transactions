-- +goose Up
-- +goose StatementBegin
ALTER TABLE transactions
    ALTER COLUMN note SET DEFAULT '',
    ALTER COLUMN note SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE transactions
    ALTER COLUMN note DROP DEFAULT,
    ALTER COLUMN note DROP NOT NULL;
-- +goose StatementEnd