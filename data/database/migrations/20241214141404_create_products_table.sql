-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100)                    NOT NULL,
    slug        VARCHAR(100)                    NOT NULL UNIQUE,
    image       TEXT                            NULL,
    rate        SMALLINT                 DEFAULT 0,
    store_id    INT REFERENCES stores           NOT NULL,
    category_id INT REFERENCES store_categories NOT NULL,
    status      SMALLINT                 DEFAULT 1,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
