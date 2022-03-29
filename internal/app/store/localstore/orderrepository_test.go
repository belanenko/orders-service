package localstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/belanenko/orders-service/internal/app/model"
	"github.com/belanenko/orders-service/internal/app/store/localstore"
)

func TestOrderRepository_SetAndGet(t *testing.T) {
	ls := localstore.New()
	m := &model.Order{
		OrderUID: "123",
	}
	ls.Order().Set(m)

	actual, _ := ls.Order().Get(m.OrderUID)

	assert.EqualValues(t, m.OrderUID, actual.OrderUID)
}

func TestOrderRepository_GetAll(t *testing.T) {
	ls := localstore.New()
	arr := []model.Order{
		{OrderUID: "1"},
		{OrderUID: "2"},
		{OrderUID: "2"},
		{OrderUID: "3"},
	}

	for _, v := range arr {
		ls.Order().Set(&v)
	}
	values, _ := ls.Order().GetAll()

	assert.Len(t, values, 3)
}
