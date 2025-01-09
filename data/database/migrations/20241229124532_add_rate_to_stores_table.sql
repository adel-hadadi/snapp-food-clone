-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS stores ADD COLUMN rate FLOAT DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS stores DROP COLUMN rate;
-- +goose StatementEnd
