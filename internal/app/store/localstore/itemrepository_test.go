package localstore_test

import (
	"testing"

	"github.com/belanenko/orders-service/internal/app/model"
	"github.com/belanenko/orders-service/internal/app/store/localstore"
	"github.com/stretchr/testify/assert"
)

func TestItemRepository_GetAndSet(t *testing.T) {
	store := localstore.New()
	key := "key"
	var value model.ItemInterface = &model.Order{}

	store.Item().Set(key, value)
	actual, _ := store.Item().Get(key)

	assert.EqualValues(t, value, actual)
}

func TestItemRepository_GetAll(t *testing.T) {
	store := localstore.New()
	items := map[string]model.ItemInterface{
		"1": &model.Order{},
		"2": &model.Order{},
		"3": &model.Order{},
	}
	for k, v := range items {
		store.Item().Set(k, v)
	}

	actual, _ := store.Item().GetAll()

	assert.EqualValues(t, items, actual)

}
