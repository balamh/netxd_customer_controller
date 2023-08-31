package main

import (
	"context"
	"fmt"
	"net"

	"github.com/balamh/netxd_customer_controller/controllers"
	"github.com/balamh/netxd_dal/netxd_dal_services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	pro "github.com/balamh/project1/netxd_customer"
)

func InitDatabase(client *mongo.Client) {
	CustomerCollection := config.GetCollection(client, "customerDB", "service")
	controllers.CustomerService = netxd_dal_services.InitCustomerService(CustomerCollection, context.Background())
}
func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	InitDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterCustomerServiceServer(s, &controllers.RPCServer{})
	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
