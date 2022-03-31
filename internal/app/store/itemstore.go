package store

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
