package sqlstore

import (
	"fmt"

	"github.com/belanenko/orders-service/internal/app/model"
)

type ItemRepository struct {
	store *Store
}

func (r *ItemRepository) Get(key string) (model.ItemInterface, error) {
	q := fmt.Sprintf("SELECT value FROM items WHERE key = '%s';", key)
	var item model.Order
	if err := r.store.db.QueryRow(q).Scan(&item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ItemRepository) Set(key string, item model.ItemInterface) error {
	value, err := item.Value()
	if err != nil {
		return err
	}
	q := fmt.Sprintf(
		"INSERT INTO items (key, value) VALUES ('%s', '%s') ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value;",
		key,
		value,
	)
	if _, err := r.store.db.Exec(q); err != nil {
		return err
	}
	return nil
}

func (r *ItemRepository) GetAll() (map[string]model.ItemInterface, error) {
	q := fmt.Sprintf("SELECT key, value FROM items;")
	var items map[string]model.ItemInterface = make(map[string]model.ItemInterface)
	rows, err := r.store.db.Query(q)
	if err != nil {
		return nil, err
	}

	var key string
	var item model.Order

	for rows.Next() {
		rows.Scan(&key, &item)

		items[key] = &item
	}

	return items, nil
}
