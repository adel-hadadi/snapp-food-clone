-- +goose Up
-- +goose StatementBegin
ALTER TABLE stores
    ADD COLUMN city_id INT REFERENCES cities;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stores
    DROP COLUMN city_id;
-- +goose StatementEnd
