-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN last_name VARCHAR NULL;

ALTER TABLE users 
    RENAME COLUMN name TO first_name;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
