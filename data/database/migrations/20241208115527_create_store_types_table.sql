-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS store_types (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    image VARCHAR(190) NOT NULL,
    url VARCHAR(100) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS store_types;
-- +goose StatementEnd
