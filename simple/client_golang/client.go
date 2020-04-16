package main

import (
	"context"
	"fmt"
	pingPongProto "github.com/juricaKenda/gRPC-PoC/simple/pb"
	"google.golang.org/grpc"
)

func Chat(stream pingPongProto.PingPongService_ChatClient) error {
	fmt.Println("Client booting..")
	err := stream.Send(message("golang_client"))
	if err != nil {
		panicOnError(err, "Client failed to notify the server about the established connection")
	}
	for {
		pong, err := stream.Recv()
		panicOnError(err, "Client failed to receive a message..")
		fmt.Println("Client received: " + pong.String())
		if pong.Message == "Current temperature is : 0" {
			askIfSnowing(stream)
		}
	}
}

func askIfSnowing(stream pingPongProto.PingPongService_ChatClient) {
	err := stream.Send(message("Wow is it snowing?"))
	panicOnError(err, "Client failed to send a message..")
}

func message(message string) *pingPongProto.Ping {
	return &pingPongProto.Ping{
		Message: message,
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

	pingerClient := pingPongProto.NewPingPongServiceClient(serverConnection)

	client, err := pingerClient.Chat(context.Background())
	panicOnError(err, "Client stub issue occurred..")

	_ = Chat(client)
}
