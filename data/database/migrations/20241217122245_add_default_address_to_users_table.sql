-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN default_address_id INT REFERENCES user_addresses NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN default_address_id;
-- +goose StatementEnd
