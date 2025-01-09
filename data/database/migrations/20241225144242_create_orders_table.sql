-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    amount INT NOT NULL,
    user_id INT REFERENCES users NOT NULL,
    user_address_id INT REFERENCES user_addresses NOT NULL,
    store_id INT REFERENCES stores NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_items (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT REFERENCES orders NOT NULL,
    product_id INT REFERENCES products NOT NULL,
    quantity SMALLINT DEFAULT 1
);

CREATE TABLE IF NOT EXISTS order_statuses (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT REFERENCES orders NOT NULL,
    status SMALLINT NOT NULL,
    user_id INT REFERENCES users NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd