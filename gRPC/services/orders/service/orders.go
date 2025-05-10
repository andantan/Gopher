package service

import (
	"context"
	"opet/gRPC/services/common/genproto/orders"
)

var (
	ordersDB = make([]*orders.Order, 0)
)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, o *orders.Order) error {
	ordersDB = append(ordersDB, o)

	return nil
}
