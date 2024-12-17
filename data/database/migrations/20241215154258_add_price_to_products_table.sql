-- +goose Up
-- +goose StatementBegin
ALTER TABLE products
    ADD COLUMN price INT NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products
    DROP COLUMN price;
-- +goose StatementEnd
