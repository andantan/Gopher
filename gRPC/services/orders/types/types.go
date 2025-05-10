package types

import (
	"context"
	"opet/gRPC/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
