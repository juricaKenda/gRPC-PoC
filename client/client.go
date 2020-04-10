package main

import (
	"context"
	"fmt"
	pingPongProto "github.com/juricaKenda/gRPC-PoC/pb"
	"google.golang.org/grpc"
	"time"
)

func Chat(stream pingPongProto.PingPongService_ChatClient) error {
	fmt.Println("Client booting..")
	for {
		pingMessage := buildPingMessage()
		fmt.Println("Client sending: " + pingMessage.String())
		err := stream.Send(pingMessage)
		panicOnError(err, "Client failed to send a message..")

		pong, err := stream.Recv()
		panicOnError(err, "Client failed to receive a message..")
		fmt.Println("Client received: " + pong.String())

		time.Sleep(time.Second)
	}
}

func buildPingMessage() *pingPongProto.Ping {
	return &pingPongProto.Ping{
		Message: "What time is it?",
	}
}

func panicOnError(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}

func closeConnection(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		fmt.Println("Issue while closing the gRPC connection..")
		panic(err)
	}
	fmt.Println("Successfully closed the gRPC connection")
}

func main() {
	serverConnection, err := grpc.Dial(":50000", grpc.WithInsecure())
	defer closeConnection(serverConnection)
	panicOnError(err, "Could not establish gRPC server connection...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pingerClient := pingPongProto.NewPingPongServiceClient(serverConnection)

	client, err := pingerClient.Chat(ctx)
	panicOnError(err, "Client stub issue occurred..")

	_ = Chat(client)
}
