-- +goose Up
CREATE TABLE source_of_incomes
(
    id         SERIAL PRIMARY KEY,
    name       TEXT      NOT NULL,
    datetime   TIMESTAMP NOT NULL,
    total      DECIMAL(10, 2),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE source_of_incomes;
