package main

import (
	"fmt"
	pingPongProto "github.com/juricaKenda/gRPC-PoC/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type Server struct {
}

func (s Server) Chat(stream pingPongProto.PingPongService_ChatServer) error {
	fmt.Println("Booting server..")
	for {
		ping, err := stream.Recv()
		panicOnError(err, "Server failed to receive a message..")

		fmt.Println("Server received: " + ping.Message)
		pongMessage := buildPongMessage(time.Now())
		fmt.Println("Server sending: " + pongMessage.String())

		err = stream.Send(pongMessage)
		panicOnError(err, "Server failed to send a message..")
	}
}

func buildPongMessage(time time.Time) *pingPongProto.Pong {
	return &pingPongProto.Pong{
		Message: time.String(),
	}
}

func panicOnError(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}
func main() {
	listener, err := net.Listen("tcp", ":50000")
	panicOnError(err, "Server failed to establish a connection..")

	grpcServer := grpc.NewServer()
	pingPongProto.RegisterPingPongServiceServer(grpcServer, &Server{})

	reflection.Register(grpcServer)

	err = grpcServer.Serve(listener)
	panicOnError(err, "gRPC server failed to serve..")
}
