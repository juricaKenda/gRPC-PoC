package main

import (
	"fmt"
	pingPongProto "github.com/juricaKenda/gRPC-PoC/simple/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type Server struct {
}

func (s Server) Chat(stream pingPongProto.PingPongService_ChatServer) error {
	waitForClientReadinessNotification(stream)
	go sendTemperatureUpdate(stream, time.Second)
	go sentTimeUpdates(stream, 2*time.Second)
	replyWhenAskedAboutSnow(stream)
	return nil
}

func waitForClientReadinessNotification(stream pingPongProto.PingPongService_ChatServer) {
	for {
		notification, err := stream.Recv()
		if err != nil {
			panicOnError(err, "Server failed while waiting for the notification of subscription")
		}
		fmt.Println(fmt.Sprintf("Client %s said he's ready! Sending temperatures and times..", notification.Message))
		break
	}
}

func sentTimeUpdates(stream pingPongProto.PingPongService_ChatServer, duration time.Duration) {
	for {
		err := stream.Send(timeUpdateMessage())
		panicOnError(err, "Server failed to send a message..")
		time.Sleep(duration)
	}
}

func replyWhenAskedAboutSnow(stream pingPongProto.PingPongService_ChatServer) {
	for {
		temperatureQuestion, err := stream.Recv()
		panicOnError(err, "Server failed to receive a message..")

		fmt.Println("Server received: " + temperatureQuestion.Message)
		err = stream.Send(&pingPongProto.Pong{
			Message: "Yes, it is snowing.",
		})
		panicOnError(err, "Server failed to send a message..")
	}
}

func sendTemperatureUpdate(stream pingPongProto.PingPongService_ChatServer, duration time.Duration) {
	temperatures := []int{10, -2, 3, 0}
	currentTempIndex := 0
	for {
		T := temperatures[currentTempIndex%len(temperatures)]
		currentTempIndex++
		err := stream.Send(temperatureUpdateMessage(T))
		panicOnError(err, "Server failed to send a message..")
		time.Sleep(duration)
	}
}

func temperatureUpdateMessage(temperature int) *pingPongProto.Pong {
	return &pingPongProto.Pong{
		Message: fmt.Sprintf("Current temperature is : %d", temperature),
	}
}

func timeUpdateMessage() *pingPongProto.Pong {
	return &pingPongProto.Pong{
		Message: time.Now().String(),
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
