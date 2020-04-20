package support

import (
	"github.com/juricaKenda/gRPC-PoC/topicbroker/infrastructure/publishers"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/infrastructure/pubsub"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/infrastructure/receiver"
	proto "github.com/juricaKenda/gRPC-PoC/topicbroker/pb"
	"google.golang.org/grpc"
	"net"
)

type TopicBroker struct {
	timepub pubsub.Publisher
	numpub  pubsub.Publisher
}

func NewTopicBroker() *TopicBroker {
	broker := new(TopicBroker)
	broker.timepub = publishers.NewTimePublisher()
	broker.timepub.Start()

	broker.numpub = publishers.NewNumberPublisher()
	broker.numpub.Start()

	return broker
}

func (tb *TopicBroker) Start() {
	lis, err := net.Listen("tcp", ":50000")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	proto.RegisterTopicServiceServer(server, NewTopicBroker())
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func (tb *TopicBroker) Pull(message *proto.Message, stream proto.TopicService_PullServer) error {
	rcvr := receiver.NewMessageReceiver()
	publisher := tb.selectPublisherByTag(message)
	publisher.Subscribe(rcvr)

	for {
		update := <-rcvr.Next()
		tb.forwardToClient(stream, update)
	}

}

func (tb *TopicBroker) receiveMessage(stream proto.TopicService_PullServer, message proto.Message) {
	err := stream.RecvMsg(message)
	if err != nil {
		panic(err)
	}
}

func (tb *TopicBroker) forwardToClient(stream proto.TopicService_PullServer, update string) {
	err := stream.Send(wrapMessage(update))
	if err != nil {
		panic(err)
	}
}

func wrapMessage(update string) *proto.Message {
	return &proto.Message{Body: update}
}

func (tb *TopicBroker) selectPublisherByTag(msg *proto.Message) pubsub.Publisher {
	switch msg.Body {
	case "time":
		return tb.timepub
	case "num":
		return tb.numpub
	}
	panic("unknown publisher tag")
}
