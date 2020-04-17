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
	broker.timeChan = make(chan string)
	broker.numChan = make(chan string)
	return broker
}

func (tb *TopicBroker) Start() {
	fmt.Println("Topic broker starting..")
}

func (tb *TopicBroker) Receive(message, pubTag string) {
	switch pubTag {
	case "time_publisher":
		tb.timeChan <- message
	case "num_publisher":
		tb.numChan <- message
	}
}
