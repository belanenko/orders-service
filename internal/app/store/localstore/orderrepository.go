package localstore

import (
	"sync"

	"github.com/belanenko/orders-service/internal/app/model"
)

type OrderRepository struct {
	m     sync.RWMutex
	store *Store
	items map[string]*model.Order
}

func (r *OrderRepository) Get(key string) (*model.Order, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	return r.items[key], nil
}

func (r *OrderRepository) Set(order *model.Order) error {
	r.m.Lock()
	defer r.m.Unlock()

	r.items[order.OrderUID] = order

	return nil
}

func (r *OrderRepository) GetAll() (map[string]*model.Order, error) {
	return r.items, nil
}
