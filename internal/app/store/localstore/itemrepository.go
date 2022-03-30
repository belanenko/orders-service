package localstore

import "github.com/belanenko/orders-service/internal/app/model"

type ItemRepository struct {
	store *Store
}

func (r *ItemRepository) Get(key string) (model.ItemInterface, error) {
	r.store.m.RLock()
	defer r.store.m.RUnlock()

	return r.store.items[key], nil
}

func (r *ItemRepository) Set(key string, item model.ItemInterface) error {
	r.store.m.Lock()
	defer r.store.m.Unlock()

	r.store.items[key] = item

	return nil
}

func (r *ItemRepository) GetAll() (map[string]model.ItemInterface, error) {
	r.store.m.RLock()
	defer r.store.m.RUnlock()

	return r.store.items, nil
}
