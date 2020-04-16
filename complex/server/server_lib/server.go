package server_lib

import (
	"context"
	"fmt"
	proto "github.com/juricaKenda/gRPC-PoC/complex/pb"
	"google.golang.org/grpc"
	"net"
	"time"
)

type TopicServer struct {
	ctx    *ServerCtx
	server *grpc.Server
}

func NewTopicServer() *TopicServer {
	server := new(TopicServer)
	server.ctx = newServerCtx()
	return server
}

func newServerCtx() *ServerCtx {
	ctx := new(ServerCtx)
	ctx.subscriptions = make(map[string][]string)
	return ctx
}

func (ts *TopicServer) Start() {
	lis, err := net.Listen("tcp", ":50000")
	panicIfErr(err, "Issue occurred while establishing server connection")

	ts.server = grpc.NewServer()
	proto.RegisterSubscriberServer(ts.server, ts.ctx)
	err = ts.server.Serve(lis)
	panicIfErr(err, "Issue occurred while starting the server")
}

type ServerCtx struct {
	Server        proto.SubscriberServer
	subscriptions map[string][]string
}

func (ctx *ServerCtx) RequestTopic(con context.Context, request *proto.Request) (*proto.Message, error) {
	fmt.Println(fmt.Sprintf("successfully subscribed %s to : %s", request.Listener.Id, request.Topic))
	topics := ctx.subscriptions[request.Listener.Id]
	topics = append(topics, request.Topic)
	ctx.subscriptions[request.Listener.Id] = topics
	return &proto.Message{Body: fmt.Sprintf("successfully subscribed to : %s", request.Topic)}, nil
}

func (ctx *ServerCtx) ListenTopic(ID *proto.ListenerID, stream proto.Subscriber_ListenTopicServer) error {
	clientTopics := ctx.subscriptions[ID.Id]
	fmt.Println(fmt.Sprintf("%s polling his topics..%s", ID.Id, clientTopics))
	for _, topic := range clientTopics {
		_ = stream.Send(topicUpdate(topic))
	}
	return nil
}

func topicUpdate(topic string) *proto.Message {
	switch topic {
	case "time_update":
		return &proto.Message{Body: time.Now().String()}
	case "axilis_update":
		return &proto.Message{Body: "Everyone doing great! See more: axilis.com"}
	}
	return &proto.Message{Body: "Unrecognized topic!"}
}

func panicIfErr(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}
