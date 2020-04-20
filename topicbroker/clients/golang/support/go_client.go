package support

import (
	"context"
	"fmt"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/pb"
	proto "github.com/juricaKenda/gRPC-PoC/topicbroker/pb"
)

type TopicClient struct {
	middleware topics.TopicServiceClient
	name       string
}

func NewTopicClient(middleware topics.TopicServiceClient, name string) *TopicClient {
	client := new(TopicClient)
	client.middleware = middleware
	client.name = name
	return client
}

func (c *TopicClient) Listen(topic string) {
	stream, err := c.middleware.Pull(context.Background(), c.buildRequest(topic))
	if err != nil {
		panic(err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(c.name + " received : " + msg.Body)
	}
}

func (c *TopicClient) buildRequest(topic string) *proto.Message {
	return &proto.Message{Body: topic}
}
