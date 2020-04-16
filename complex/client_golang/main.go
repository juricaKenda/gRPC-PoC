package main

import (
	"fmt"
	"github.com/juricaKenda/gRPC-PoC/complex/client_golang/client_lib"
	proto "github.com/juricaKenda/gRPC-PoC/complex/pb"
	"google.golang.org/grpc"
	"time"
)

func main() {
	connection, err := grpc.Dial(":50000", grpc.WithInsecure())
	panicIfErr(err, "Issue while establishing a connection")

	listener := client_lib.NewListenerContext(proto.NewSubscriberClient(connection), "fico")
	listener.RequestTopicAfter("time_update", 0*time.Second)
	go listener.RequestTopicAfter("axilis_update", 5*time.Second)
	listener.Listen()

}

func panicIfErr(err error, message string) {
	if err != nil {
		fmt.Println(message)
		panic(err)
	}
}
