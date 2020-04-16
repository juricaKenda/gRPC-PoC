package client_lib

import (
	"context"
	"fmt"
	proto "github.com/juricaKenda/gRPC-PoC/complex/pb"
)

type ListenerCtx struct {
	Client proto.SubscriberClient
	ID     proto.ListenerID
}

func NewListenerContext(client proto.SubscriberClient) *ListenerCtx {
	return &ListenerCtx{
		Client: client,
		ID: proto.ListenerID{
			Id: "golang_client",
		},
	}
}

func (ctx *ListenerCtx) RequestTopic(topic string) {
	confirmation, err := ctx.Client.RequestTopic(context.Background(), ctx.buildRequest(topic))
	panicIfErr(err, "Issue occurred while requesting a topic")
	fmt.Println(confirmation.Body)
}

func panicIfErr(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}

func (ctx *ListenerCtx) buildRequest(topic string) *proto.Request {
	return &proto.Request{
		Listener: &ctx.ID,
		Topic:    topic,
	}
}

func (ctx *ListenerCtx) ListenTopics() {
	stream, err := ctx.Client.ListenTopic(context.Background(), &ctx.ID)
	panicIfErr(err, "Issue occurred while listening to topics")

	for {
		updateMessage, err := stream.Recv()
		panicIfErr(err, "Issue occurred while receiving topic update")
		fmt.Println(updateMessage)
	}

}
