package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"snapp-food/internal/entity"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return ProductRepository{db: db}
}

func (r ProductRepository) Find(ctx context.Context, id int) (entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM products WHERE id=$1`

	var product entity.Product
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Slug,
		&product.Image,
		&product.Rate,
		&product.StoreID,
		&product.CategoryID,
		&product.Status,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.Price,
	)

	return product, err
}

func (r ProductRepository) FindBySlug(ctx context.Context, slug string) (entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM products WHERE slug=$1`

	var product entity.Product
	err := r.db.QueryRowContext(ctx, query, slug).Scan(
		&product.ID,
		&product.Name,
		&product.Slug,
		&product.Image,
		&product.Rate,
		&product.StoreID,
		&product.CategoryID,
		&product.Status,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.Price,
	)

	return product, err
}

func (r ProductRepository) Get(ctx context.Context) ([]entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM products`

	products := make([]entity.Product, 0)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Slug,
			&product.Image,
			&product.Rate,
			&product.StoreID,
			&product.CategoryID,
			&product.Status,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r ProductRepository) Create(ctx context.Context, store entity.Product) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO products 
		(name, slug, image, store_id, category_id, price) 
		VALUES 
		($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(
		ctx,
		query,
		store.Name,
		store.Slug,
		store.Image,
		store.StoreID,
		store.CategoryID,
		store.Price,
	)

	return err
}

func (r ProductRepository) Update(ctx context.Context, id int, product entity.Product) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE products SET name=$1, price=$2, image=$3, status=$4 WHERE id=$5`

	_, err := r.db.ExecContext(
		ctx,
		query,
		product.Name,
		product.Price,
		product.Image,
		product.Status,
		id,
	)

	return err
}

func (r ProductRepository) GetByStoreID(ctx context.Context, storeID int) ([]entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM products WHERE store_id = $1`

	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]entity.Product, 0)
	for rows.Next() {
		var product entity.Product

		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Slug,
			&product.Image,
			&product.Rate,
			&product.StoreID,
			&product.CategoryID,
			&product.Status,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Price,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
