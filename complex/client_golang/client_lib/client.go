package client_lib

import (
	"context"
	"fmt"
	proto "github.com/juricaKenda/gRPC-PoC/complex/pb"
	"io"
	"time"
)

type ListenerCtx struct {
	Client proto.SubscriberClient
	ID     proto.ListenerID
}

func NewListenerContext(client proto.SubscriberClient, name string) *ListenerCtx {
	return &ListenerCtx{
		Client: client,
		ID: proto.ListenerID{
			Id: name,
		},
	}
}

func (ctx *ListenerCtx) RequestTopicAfter(topic string, duration time.Duration) {
	time.Sleep(duration)
	confirmation, err := ctx.Client.RequestTopic(context.Background(), ctx.buildRequest(topic))
	panicIfErr(err, "Issue occurred while requesting a topic")
	fmt.Println(confirmation.Body)
}

func (ctx *ListenerCtx) Listen() {
	for {
		//ctx.Poll()
		time.Sleep(2 * time.Second)
	}
}
func isEOF(err error) bool {
	return err == io.EOF
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

func (ctx *ListenerCtx) Poll() {
	stream, err := ctx.Client.ListenTopic(context.Background(), &ctx.ID)
	panicIfErr(err, "Issue occurred while listening to topics")
	for {
		updateMessage, err := stream.Recv()
		if isEOF(err) {
			break
		}
		panicIfErr(err, "Issue occurred while receiving topic update")
		fmt.Println(updateMessage)
	}

}
