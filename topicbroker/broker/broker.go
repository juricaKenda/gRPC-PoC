package broker

import (
	"fmt"
)

type TopicBroker struct {
	timeChan chan string
	numChan  chan string
}

func NewTopicBroker() *TopicBroker {
	broker := new(TopicBroker)

	return broker
}

func (tb *TopicBroker) Start() {
	fmt.Println("Topic broker starting..")
}

func (tb *TopicBroker) Receive(message, pubTag string) {
	panic("implement me")
}
