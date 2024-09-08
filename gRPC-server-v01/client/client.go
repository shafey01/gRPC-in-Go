package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/shafey01/gRPC-in-Go/gRPC-server-v01/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	connection, err := grpc.NewClient("localhost:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to grpc server %v", err)
	}
	defer connection.Close()

	client := pb.NewCoffeShopClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menuStream, err := client.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatalf("error calling func menuStream %v", err)

	}

	done := make(chan bool)
	var items []*pb.Item

	go func() {
		for {
			response, err := menuStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("can not receive from getmenu func %v", err)
			}
			items = response.Items
			log.Printf("Response received: %v", response.Items)
		}
	}()
	var isDone bool
	isDone = <-done
	log.Printf("log from Done channel %v", isDone)

	receipt, err := client.PlaceOrder(ctx, &pb.Order{Items: items})
	log.Printf("Receipt %v", receipt)

	status, err := client.GetOrderStatus(ctx, receipt)
	log.Printf("order status %v", status)

}
