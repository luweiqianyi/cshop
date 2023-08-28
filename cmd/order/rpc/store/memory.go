package store

import (
	"cshop/cmd/order/rpc/store/entity"
	"fmt"
	"sync"
)

var memStoreOnce sync.Once
var gMemStore *MemoryStore

func MemStoreInstance() *MemoryStore {
	memStoreOnce.Do(func() {
		gMemStore = &MemoryStore{}
	})
	return gMemStore
}

type MemoryStore struct {
	orderMap sync.Map
}

func (m *MemoryStore) Add(orderID string, order entity.Order) error {
	_, found := m.orderMap.Load(orderID)
	if found {
		return fmt.Errorf("add: order[%v] already exist", orderID)
	}
	m.orderMap.Store(orderID, order)
	return nil
}

func (m *MemoryStore) Delete(orderID string) error {
	m.orderMap.Delete(orderID)
	return nil
}

func (m *MemoryStore) Query(orderID string) (entity.Order, error) {
	order, found := m.orderMap.Load(orderID)
	if !found {
		return entity.Order{}, fmt.Errorf("query: order[%v] not exist", orderID)
	}

	ret, ok := order.(entity.Order)
	if !ok {
		return entity.Order{}, fmt.Errorf("query: order[%v] format error", orderID)
	}
	return ret, nil
}

func (m *MemoryStore) UpdateOrderState(orderID string, status int) error {
	order, found := m.orderMap.Load(orderID)
	if !found {
		return fmt.Errorf("update order status: order[%v] not exist", orderID)
	}

	ret, ok := order.(entity.Order)
	if !ok {
		return fmt.Errorf("update order status: order[%v] format error", orderID)
	}

	ret.OrderStatus = status

	m.orderMap.Store(orderID, ret)

	return nil
}
