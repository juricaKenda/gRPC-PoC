package publishers

import (
	"fmt"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/pubsub"
	"time"
)

type numpub struct {
	subscribers []pubsub.Subscriber
	tag         string
}

func NewNumberPublisher() pubsub.Publisher {
	return &numpub{
		subscribers: make([]pubsub.Subscriber, 0),
		tag:         "number_publisher",
	}
}

func (n *numpub) Start() {
	fmt.Println("Starting number publisher...")
	go n.run()
}

func (n *numpub) NotifyAll() {
	message := "123"
	for _, sub := range n.subscribers {
		sub.Receive(message, n.tag)
	}
}

func (n *numpub) Subscribe(subscriber pubsub.Subscriber) {
	n.subscribers = append(n.subscribers, subscriber)
}

func (n *numpub) run() {
	for {
		n.NotifyAll()
		time.Sleep(2 * time.Second)
	}
}
