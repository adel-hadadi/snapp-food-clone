package repository

import (
	"context"
	"time"

	"snapp-food/internal/entity"

	"github.com/jmoiron/sqlx"
)

type StoreCategoryRepository struct {
	db *sqlx.DB
}

func NewStoreCategoryRepository(db *sqlx.DB) StoreCategoryRepository {
	return StoreCategoryRepository{db: db}
}

func (r StoreCategoryRepository) Find(ctx context.Context, id int) (entity.StoreCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM store_categories WHERE id=$1`

	var storeCategory entity.StoreCategory
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&storeCategory.ID,
		&storeCategory.Name,
		&storeCategory.StoreID,
	)

	return storeCategory, err
}

func (r StoreCategoryRepository) Get(ctx context.Context) ([]entity.StoreCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM store_categories`

	storeCategories := make([]entity.StoreCategory, 0)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var storeCategory entity.StoreCategory

		err := rows.Scan(
			&storeCategory.ID,
			&storeCategory.Name,
			&storeCategory.StoreID,
		)
		if err != nil {
			return nil, err
		}

		storeCategories = append(storeCategories, storeCategory)
	}

	return storeCategories, nil
}

func (r StoreCategoryRepository) GetByStoreID(ctx context.Context, storeID int) ([]entity.StoreCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM store_categories WHERE store_id=$1`

	rows, err := r.db.QueryContext(ctx, query, storeID)
	if err != nil {
		return nil, err
	}

	categories := make([]entity.StoreCategory, 0)
	for rows.Next() {
		var c entity.StoreCategory

		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.StoreID,
		); err != nil {
			return nil, err
		}

		query = `SELECT id, name, slug, image, rate, price, description FROM products WHERE category_id = $1`

		productRows, err := r.db.QueryContext(ctx, query, c.ID)
		if err != nil {
			return nil, err
		}

		for productRows.Next() {
			var p entity.Product
			if err := productRows.Scan(
				&p.ID,
				&p.Name,
				&p.Slug,
				&p.Image,
				&p.Rate,
				&p.Price,
				&p.Description,
			); err != nil {
				return nil, err
			}

			c.Products = append(c.Products, p)
		}

		categories = append(categories, c)
	}

	return categories, nil
}

func (r StoreCategoryRepository) Create(ctx context.Context, store entity.StoreCategory) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO store_categories 
		(name, store_id) 
		VALUES 
		($1, $2)`

	_, err := r.db.ExecContext(
		ctx,
		query,
		store.Name,
		store.StoreID,
	)

	return err
}

func (r StoreCategoryRepository) Update(ctx context.Context, id int, product entity.StoreCategory) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE store_categories SET name=$1  WHERE id=$5`

	_, err := r.db.ExecContext(
		ctx,
		query,
		product.Name,
		id,
	)

	return err
}
