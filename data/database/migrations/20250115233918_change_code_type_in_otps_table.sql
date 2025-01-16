-- +goose Up
-- +goose StatementBegin
ALTER TABLE otps ALTER COLUMN code TYPE VARCHAR(10);
-- +goose StatementEnd
