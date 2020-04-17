package broker

import (
	"fmt"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/publishers"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/pubsub"
)

type TopicBroker struct {
	timepub pubsub.Publisher
	numpub  pubsub.Publisher
}
type Receiver struct {
	messages chan string
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

func (r *Receiver) Receive(message, pubTag string) {
	r.messages <- message
}

func (tb *TopicBroker) Test() {
	receiver := &Receiver{
		messages: make(chan string),
	}
	tb.numpub.Subscribe(receiver)

	for {
		message := <-receiver.messages
		fmt.Println(fmt.Sprintf("Got a message: %s", message))
	}

}
