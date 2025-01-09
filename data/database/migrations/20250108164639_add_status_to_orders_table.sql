-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders ADD COLUMN status SMALLINT DEFAULT 0;
-- +goose StatementEnd
