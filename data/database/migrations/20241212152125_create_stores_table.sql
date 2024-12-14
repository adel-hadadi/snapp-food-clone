-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stores
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(100)               NOT NULL,
    slug          VARCHAR(100)               NOT NULL UNIQUE,
    latitude      FLOAT                      NOT NULL,
    longitude     FLOAT                      NOT NULL,
    logo          TEXT                       NULL,
    manager_id    INT REFERENCES users       NOT NULL,
    store_type_id INT REFERENCES store_types NOT NULL,
    status        SMALLINT                 DEFAULT 1,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stores;
-- +goose StatementEnd
