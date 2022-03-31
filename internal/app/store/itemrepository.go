package store

import "github.com/belanenko/orders-service/internal/app/model"

type ItemRepository struct {
	store *Store
}

func (r *ItemRepository) Get(key string) (model.ItemInterface, error) {
	return r.store.localstore.Item().Get(key)
}

func (r *ItemRepository) Set(key string, value model.ItemInterface) error {
	if err := r.store.localstore.Item().Set(key, value); err != nil {
		return nil
	}

	if err := r.store.sqlstore.Item().Set(key, value); err != nil {
		return err
	}

	return nil
}

func (r *ItemRepository) GetAll() (map[string]model.ItemInterface, error) {
	return r.store.localstore.Item().GetAll()
}
