package repository

import (
	"context"
	"time"

	"snapp-food/internal/entity"

	"github.com/jmoiron/sqlx"
)

type ProductCategoryRepository struct {
	db *sqlx.DB
}

func NewProductCategoryRepository(db *sqlx.DB) ProductCategoryRepository {
	return ProductCategoryRepository{db: db}
}

func (r ProductCategoryRepository) Find(ctx context.Context, id int) (entity.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM product_categories WHERE id=$1`

	var productCategory entity.ProductCategory
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&productCategory.ID,
		&productCategory.Name,
	)

	return productCategory, err
}

func (r ProductCategoryRepository) Get(ctx context.Context) ([]entity.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM product_categories`

	productCategories := make([]entity.ProductCategory, 0)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productCategory entity.ProductCategory

		err := rows.Scan(
			&productCategory.ID,
			&productCategory.Name,
			&productCategory.Slug,
			&productCategory.Image,
		)
		if err != nil {
			return nil, err
		}

		productCategories = append(productCategories, productCategory)
	}

	return productCategories, nil
}

func (r ProductCategoryRepository) Create(ctx context.Context, category entity.ProductCategory) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO product_categories 
		(name, slug, image)
		VALUES
		($1, $2)`

	_, err := r.db.ExecContext(
		ctx,
		query,
		category.Name,
		category.Slug,
		category.Image,
	)

	return err
}

func (r ProductCategoryRepository) Update(ctx context.Context, id int, product entity.ProductCategory) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE product_categories SET name=$1  WHERE id=$5`

	_, err := r.db.ExecContext(
		ctx,
		query,
		product.Name,
		id,
	)

	return err
}
