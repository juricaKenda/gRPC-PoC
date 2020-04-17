package main

import (
	"github.com/juricaKenda/gRPC-PoC/topicbroker/broker"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/publishers"
)

func main() {
	topicBroker := broker.NewTopicBroker()

	timepub := publishers.NewTimePublisher()
	timepub.Subscribe(topicBroker)
	timepub.Start()

	numpub := publishers.NewNumberPublisher()
	numpub.Subscribe(topicBroker)
	numpub.Start()

	topicBroker.Start()
}
