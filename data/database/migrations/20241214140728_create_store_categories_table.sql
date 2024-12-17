-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS store_categories
(
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(50) NOT NULL,
    store_id INT REFERENCES stores
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS store_categories;
-- +goose StatementEnd
