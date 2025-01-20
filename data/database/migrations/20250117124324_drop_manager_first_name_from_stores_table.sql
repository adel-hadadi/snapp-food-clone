-- +goose Up
-- +goose StatementBegin
ALTER TABLE stores DROP COLUMN manager_first_name;
ALTER TABLE stores DROP COLUMN manager_last_name;
ALTER TABLE stores DROP COLUMN phone;
-- +goose StatementEnd
