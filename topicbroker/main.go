package main

import (
	"github.com/juricaKenda/gRPC-PoC/topicbroker/broker"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/publishers"
)

func main() {

	timepub := publishers.NewTimePublisher()
	timepub.Start()

	numpub := publishers.NewNumberPublisher()
	numpub.Start()
	topicBroker := broker.NewTopicBroker()
	topicBroker.Start()
}
