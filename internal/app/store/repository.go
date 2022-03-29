package store

import "github.com/belanenko/orders-service/internal/app/model"

type OrderRepository interface {
	Get(string) (*model.Order, error)
	Set(*model.Order) error

	GetAll() (map[string]*model.Order, error)
}
