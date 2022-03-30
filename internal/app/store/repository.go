package store

import "github.com/belanenko/orders-service/internal/app/model"

type ItemRepositoryInterface interface {
	Get(key string) (model.ItemInterface, error)
	Set(key string, item model.ItemInterface) error

	GetAll() (map[string]model.ItemInterface, error)
}
