package store

import "cshop/cmd/order/rpc/store/entity"

// OrderStore 用户-订单接口(一个用户可以创建多个订单，一个订单只能属于一个用户)
type OrderStore interface {
	Add(orderID string, order entity.Order) error
	Delete(orderID string) error
	Query(orderID string) (entity.Order, error)
	UpdateOrderState(orderID string, status int) error
}
