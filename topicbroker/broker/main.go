package main

import "github.com/juricaKenda/gRPC-PoC/topicbroker/broker/support"

func main() {
	broker := support.NewTopicBroker()
	broker.Start()
}
