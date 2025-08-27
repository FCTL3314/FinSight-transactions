-- +goose Up
-- +goose StatementBegin
ALTER TABLE finance_detailing
    ALTER COLUMN date_from TYPE DATE USING date_from::date,
    ALTER COLUMN date_to TYPE DATE USING date_to::date;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE finance_detailing
    ALTER COLUMN date_from TYPE TIMESTAMPTZ USING date_from::timestamptz,
    ALTER COLUMN date_to TYPE TIMESTAMPTZ USING date_to::timestamptz;
-- +goose StatementEnd