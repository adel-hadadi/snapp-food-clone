-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_addresses
(
    id          SERIAL PRIMARY KEY NOT NULL,
    name        VARCHAR(90)        NOT NULL,
    latitude    FLOAT              NOT NULL,
    longitude   FLOAT              NOT NULL,
    user_id     INT REFERENCES users,
    city_id     INT REFERENCES cities,
    province_id INT REFERENCES provinces,
    address     TEXT               NOT NULL,
    detail      VARCHAR(100)       NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_addresses;
-- +goose StatementEnd
