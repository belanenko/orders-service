package sqlstore

import (
	"database/sql"

	"github.com/belanenko/orders-service/internal/app/store"
)

type Store struct {
	db             *sql.DB
	itemRepository store.ItemRepositoryInterface
}

func (s *Store) Item() store.ItemRepositoryInterface {
	if s.itemRepository != nil {
		return s.itemRepository
	}

	s.itemRepository = &ItemRepository{
		store: s,
	}
	return s.itemRepository
}

func New(db *sql.DB) store.StoreInterface {
	return &Store{
		db: db,
	}
}
