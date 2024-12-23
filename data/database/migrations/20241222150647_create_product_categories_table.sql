-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(100) NOT NULL UNIQUE,
    image VARCHAR NOT NULL
);

ALTER TABLE products ADD COLUMN product_category_id INT REFERENCES product_categories;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product_categories;
-- +goose StatementEnd
