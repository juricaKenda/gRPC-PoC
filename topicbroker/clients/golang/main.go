package main

import (
	"github.com/juricaKenda/gRPC-PoC/topicbroker/clients/golang/support"
	proto "github.com/juricaKenda/gRPC-PoC/topicbroker/pb"
	"google.golang.org/grpc"
)

func main() {
	fico := startTopicClient(":50000", "fico")
	go fico.Listen("num")
	go fico.Listen("time")

	kenda := startTopicClient(":50000", "kenda")
	go kenda.Listen("time")
	select {}
}

func startTopicClient(target, name string) *support.TopicClient {
	connection, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	serviceClient := proto.NewTopicServiceClient(connection)
	client := support.NewTopicClient(serviceClient, name)
	return client
}
