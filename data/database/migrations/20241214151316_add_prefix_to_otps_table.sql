-- +goose Up
-- +goose StatementBegin
ALTER TABLE otps
    ADD COLUMN prefix VARCHAR(20) NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE otps
    DROP COLUMN prefix;
-- +goose StatementEnd
