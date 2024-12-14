-- +goose Up
-- +goose StatementBegin
ALTER TABLE stores
    DROP CONSTRAINT stores_manager_id_fkey;
ALTER TABLE stores
    DROP COLUMN manager_id;

ALTER TABLE stores
    ADD COLUMN manager_first_name VARCHAR(50) NOT NULL;
ALTER TABLE stores
    ADD COLUMN manager_last_name VARCHAR(50) NOT NULL;
ALTER TABLE stores
    ADD COLUMN phone VARCHAR(10) NOT NULL UNIQUE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
