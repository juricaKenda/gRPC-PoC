package broker

import (
	"fmt"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/publishers"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/pubsub"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/receiver"
)

type TopicBroker struct {
	timepub pubsub.Publisher
	numpub  pubsub.Publisher
}

func NewTopicBroker() *TopicBroker {
	broker := new(TopicBroker)
	broker.timepub = publishers.NewTimePublisher()
	broker.timepub.Start()

	broker.numpub = publishers.NewNumberPublisher()
	broker.numpub.Start()
	return broker
}

func (tb *TopicBroker) Start() {
	fmt.Println("Topic broker starting..")
	tb.Test()
}

func (tb *TopicBroker) Test() {
	rcvr := receiver.NewMessageReceiver()
	tb.numpub.Subscribe(rcvr)

	for {
		message := <-rcvr.Next()
		fmt.Println(fmt.Sprintf("Got a message: %s", message))
	}

}
