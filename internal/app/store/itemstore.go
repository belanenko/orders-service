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

	// rows, err := s.localstore.Item().GetAll()
	// if err != nil {
	// 	log.Fatal("-_- скачать кеш с бд не вышло, я упал")
	// }

	// for _, row := range rows {
	// 	d, _ := row.Value()
	// 	s.localstore.Item().Set(d row)
	// }

	s.itemRepository = &ItemRepository{
		store: s,
	}

	return s.itemRepository
}

func New(localstore StoreInterface, sqlstore StoreInterface) *Store {
	return &Store{
		localstore: localstore,
		sqlstore:   sqlstore,
	}
}
