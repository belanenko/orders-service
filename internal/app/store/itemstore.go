package store

import (
	"encoding/json"

	"github.com/belanenko/orders-service/internal/app/model"
)

type Store struct {
	localstore     StoreInterface
	sqlstore       StoreInterface
	itemRepository ItemRepositoryInterface
}

func (s *Store) Item() ItemRepositoryInterface {
	if s.itemRepository != nil {
		return s.itemRepository
	}

	s.itemRepository = &ItemRepository{
		store: s,
	}

	return s.itemRepository
}

func (s Store) LocalCache() error {
	rows, err := s.sqlstore.Item().GetAll()
	if err != nil {
		return err
	}

	for _, row := range rows {
		valueJson := row.Json()
		var item model.Order
		if err := json.Unmarshal(valueJson, &item); err != nil {
			return err
		}
		if err := s.localstore.Item().Set(item.OrderUID, &item); err != nil {
			return err
		}
	}

	return nil
}

func New(localstore StoreInterface, sqlstore StoreInterface) *Store {
	s := &Store{
		localstore: localstore,
		sqlstore:   sqlstore,
	}

	s.LocalCache()

	return s
}
