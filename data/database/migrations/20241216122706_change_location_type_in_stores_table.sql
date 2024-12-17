-- +goose Up
-- +goose StatementBegin
ALTER TABLE stores DROP COLUMN latitude;
ALTER TABLE stores DROP COLUMN longitude;

ALTER TABLE stores ADD COLUMN location geography;

ALTER TABLE user_addresses DROP COLUMN latitude;
ALTER TABLE user_addresses DROP COLUMN longitude;

ALTER TABLE user_addresses ADD COLUMN location geography;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
