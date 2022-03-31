package sqlstore

import (
	"testing"

	"github.com/belanenko/orders-service/internal/app/model"
	_ "github.com/lib/pq" // ...
	"github.com/stretchr/testify/assert"
)

func TestItemRepository_Set(t *testing.T) {
	db, teardown := TestDB(t, "host=127.0.0.1 port=5432 user=rwuser password=password dbname=dev sslmode=disable")
	defer teardown("items")

	key, value := "mew1", &model.Order{OrderUID: "qwe"}
	storeDb := New(db)
	err := storeDb.Item().Set(key, value)
	assert.NoError(t, err)

}

func TestItemRepository_Get(t *testing.T) {
	db, teardown := TestDB(t, "host=127.0.0.1 port=5432 user=rwuser password=password dbname=dev sslmode=disable")
	defer teardown("items")

	key, value := "mew", &model.Order{OrderUID: "mewmwemw"}

	storeDb := New(db)
	err := storeDb.Item().Set(key, value)
	assert.NoError(t, err)
	actual, err := storeDb.Item().Get(key)
	assert.NoError(t, err)

	var expect model.ItemInterface = value

	assert.EqualValues(t, expect, actual)

}

func TestItemRepository_GetAll(t *testing.T) {
	db, teardown := TestDB(t, "host=127.0.0.1 port=5432 user=rwuser password=password dbname=dev sslmode=disable")
	defer teardown("items")

	items := make(map[string]model.ItemInterface)
	items["1"] = &model.Order{}
	items["2"] = &model.Order{}
	items["3"] = &model.Order{}

	storeDb := New(db)

	for k, v := range items {
		storeDb.Item().Set(k, v)
	}

	actual, err := storeDb.Item().GetAll()
	assert.NoError(t, err)
	assert.EqualValues(t, len(items), len(actual))
	assert.EqualValues(t, items["1"], actual["1"])

}
