package main

import (
	"github.com/juricaKenda/gRPC-PoC/topicbroker/clients/golang/support"
	proto "github.com/juricaKenda/gRPC-PoC/topicbroker/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial(":50000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	serviceClient := proto.NewTopicServiceClient(connection)
	client := support.NewTopicClient(serviceClient, "fico")
	client.Listen("num")
}
