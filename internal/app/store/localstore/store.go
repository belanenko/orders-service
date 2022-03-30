package localstore

import (
	"sync"

	"github.com/belanenko/orders-service/internal/app/model"
	"github.com/belanenko/orders-service/internal/app/store"
)

type Store struct {
	m              sync.RWMutex
	items          map[string]model.ItemInterface
	itemRepository store.ItemRepositoryInterface
}

func (s *Store) Item() store.ItemRepositoryInterface {
	if s.itemRepository != nil {
		return s.itemRepository
	}

	s.items = make(map[string]model.ItemInterface)
	s.itemRepository = &ItemRepository{
		store: s,
	}

	return s.itemRepository
}

func New() store.StoreInterface {
	return &Store{}
}
