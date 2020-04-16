package main

import "github.com/juricaKenda/gRPC-PoC/complex/server/server_lib"

func main() {
	topicServer := server_lib.NewTopicServer()
	topicServer.Start()
}
