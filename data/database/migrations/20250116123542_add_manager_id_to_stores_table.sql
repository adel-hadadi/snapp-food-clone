-- +goose Up
-- +goose StatementBegin
ALTER TABLE stores ADD COLUMN manager_id BIGINT REFERENCES users;
-- +goose StatementEnd