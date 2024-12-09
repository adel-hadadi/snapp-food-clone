-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS provinces
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS cities
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(17) NOT NULL,
    province_id INT REFERENCES provinces
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS provinces;
DROP TABLE IF EXISTS cities;
-- +goose StatementEnd
