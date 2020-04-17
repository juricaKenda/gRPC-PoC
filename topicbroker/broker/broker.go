package broker

import (
	"fmt"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/publishers"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/pubsub"
)

type TopicBroker struct {
	timepub pubsub.Publisher
}

func NewTopicBroker() *TopicBroker {
	broker := new(TopicBroker)
	broker.timepub = publishers.NewTimePublisher()
	broker.timepub.Start()
	return broker
}

func (tb *TopicBroker) Start() {
	fmt.Println("Topic broker starting..")
}
