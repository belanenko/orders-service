package localstore

import "github.com/belanenko/orders-service/internal/app/model"

type Store struct {
	orderRepository *OrderRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Order() *OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}

	s.orderRepository = &OrderRepository{
		store: s,
	}

	if s.orderRepository.items == nil {
		s.orderRepository.items = make(map[string]*model.Order)
	}

	return s.orderRepository
}
