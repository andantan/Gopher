package handler

import (
	"context"
	"opet/gRPC/services/common/genproto/orders"
	"opet/gRPC/services/orders/types"

	"google.golang.org/grpc"
)

type OrdersGRPCHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGRPCOrdersService(grpc *grpc.Server, os types.OrderService) {
	gRPCHandler := &OrdersGRPCHandler{
		orderService: os,
	}

	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGRPCHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderId:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}

	err := h.orderService.CreateOrder(ctx, order)

	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}
