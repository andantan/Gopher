package main

import (
	"log"
	"net"
	handler "opet/gRPC/services/orders/handler/orders"
	"opet/gRPC/services/orders/service"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	// register grpc services
	orderService := service.NewOrderService()
	handler.NewGRPCOrdersService(grpcServer, orderService)
	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
