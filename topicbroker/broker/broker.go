package broker

import (
	"fmt"
)

type TopicBroker struct {
}

func NewTopicBroker() *TopicBroker {
	broker := new(TopicBroker)

	return broker
}

func (tb *TopicBroker) Start() {
	fmt.Println("Topic broker starting..")
}
