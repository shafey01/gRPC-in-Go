package main

import (
	"context"

	pb "github.com/shafey01/gRPC-in-Go/gRPC-server-v01/proto"
)

type server struct{
	pb.UnimplementedCoffeShopServer
}


func (s *server) GetMenu(menuRequest *pb.MenuRequest, svr pb.CoffeShop_GetMenuServer) error {
	return nil
}
func (s *server) PlaceOrder(context.Context, *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "A1234",

	},nil
}
func (s *server) GetOrderStatus( context.Context, *pb.Receipt) (*pb.OrderStatus, error) {

}
func (s *server) mustEmbedUnimplementedCoffeShopServer() {}
