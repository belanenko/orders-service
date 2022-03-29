package sqlstore

import (
	"fmt"

	"github.com/belanenko/orders-service/internal/app/model"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Get(key string) (*model.Order, error) {
	var order model.Order
	if err := r.store.db.QueryRow(
		"SELECT * FROM orders WHERE key = $1", key,
	).Scan(&order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) Set(order *model.Order) error {
	driverValue, err := order.Value()
	if err != nil {
		return err
	}

	q := fmt.Sprintf(
		"INSERT INTO orders (key, value) VALUES (%s, '%s')",
		order.OrderUID,
		driverValue,
	)
	if _, err := r.store.db.Exec(q); err != nil {
		return err
	}

	return nil
}
