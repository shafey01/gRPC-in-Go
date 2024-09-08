package main

import (
	"context"
	"log"
	"net"

	pb "github.com/shafey01/gRPC-in-Go/gRPC-server-v01/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeShopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest, svr pb.CoffeShop_GetMenuServer) error {
	items := []*pb.Item{
		&pb.Item{

			Id:   "1",
			Name: "Black",
		},

		&pb.Item{

			Id:   "2",
			Name: "Latee",
		},
		&pb.Item{

			Id:   "3",
			Name: "Mocha",
		},
	}

	for i, _ := range items {
		svr.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
	}
	return nil
}
func (s *server) PlaceOrder(context.Context, *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "A1234",
	}, nil
}
func (s *server) GetOrderStatus(ctx context.Context, recepit *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: recepit.Id,
		Status:  "In Progress",
	}, nil
}

func main() {

	listener, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalf("Error in establish the server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCoffeShopServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}

	log.Println("server running in port: 9002")

}
