package repository

import (
	"context"
	"snapp-food/internal/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) OrderRepository {
	return OrderRepository{
		db: db,
	}
}

func (r OrderRepository) Create(ctx context.Context, userID, userAddressID, storeID, amount int, items []entity.OrderItem) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO orders (user_id, user_address_id, store_id, amount) VALUES ($1, $2, $3, $4) RETURNING id`

	var orderID int

	if err := r.db.QueryRowContext(ctx, query, userID, userAddressID, storeID, amount).Scan(&orderID); err != nil {
		return err
	}

	query = `INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`
	for _, item := range items {
		_, err := r.db.ExecContext(ctx, query, orderID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return err
		}
	}

	// TODO: should add status created in order_statuses

	return nil
}

func (r OrderRepository) GetByID(ctx context.Context, orderID int) (entity.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM orders WHERE orders.id = $1`

	var order entity.Order
	if err := r.db.QueryRowxContext(ctx, query, orderID).StructScan(&order); err != nil {
		return entity.Order{}, err
	}

	return order, nil
}

func (r OrderRepository) GetByUserID(ctx context.Context, userID int) ([]entity.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT 
	orders.*, stores.id, stores.name, stores.slug, stores.logo
	FROM orders
	LEFT JOIN stores ON stores.id = orders.store_id
	WHERE orders.user_id = $1
    ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]entity.Order, 0)
	query = `
		SELECT order_items.*,
			products.id, products.name, products.slug, products.image
		FROM order_items 
		LEFT JOIN products ON products.id = order_items.product_id
		WHERE order_items.order_id=$1
	`

	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(
			&order.ID,
			&order.Amount,
			&order.UserID,
			&order.UserAddressID,
			&order.StoreID,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.Status,
			&order.Store.ID,
			&order.Store.Name,
			&order.Store.Slug,
			&order.Store.Logo,
		); err != nil {
			return nil, err
		}

		rows, err := r.db.QueryContext(ctx, query, order.ID)
		if err != nil {
			return nil, err
		}

		items := make([]entity.OrderItem, 0)
		for rows.Next() {
			var orderItem entity.OrderItem
			if err := rows.Scan(
				&orderItem.ID,
				&orderItem.OrderID,
				&orderItem.ProductID,
				&orderItem.Quantity,
				&orderItem.Price,
				&orderItem.Product.ID,
				&orderItem.Product.Name,
				&orderItem.Product.Slug,
				&orderItem.Product.Image,
			); err != nil {
				return nil, err
			}

			items = append(items, orderItem)
		}
		order.Items = items

		orders = append(orders, order)
		rows.Close()
	}

	return orders, nil
}

func (r OrderRepository) UpdateStatus(ctx context.Context, orderID int, status int8) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE orders SET status = $1 WHERE orders.id = $2`

	_, err := r.db.ExecContext(ctx, query, status, orderID)

	return err
}

func (r OrderRepository) RemovePending(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `DELETE FROM orders WHERE status=$1`

	_, err := r.db.ExecContext(ctx, query, entity.OrderStatusPending)
	return err
}
