package main

import (
	"github.com/juricaKenda/gRPC-PoC/topicbroker/broker"
)

func main() {
	topicBroker := broker.NewTopicBroker()
	topicBroker.Start()

}
