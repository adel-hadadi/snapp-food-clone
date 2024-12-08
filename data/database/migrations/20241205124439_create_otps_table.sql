-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS otps (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    phone VARCHAR(10) NOT NULL,
    code INT NOT NULL,
    status SMALLINT DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS otps;
-- +goose StatementEnd
