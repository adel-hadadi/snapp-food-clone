-- +goose Up
-- +goose StatementBegin
ALTER TABLE order_items
    ADD COLUMN price INT NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE order_items DROP COLUMN price;
-- +goose StatementEnd
