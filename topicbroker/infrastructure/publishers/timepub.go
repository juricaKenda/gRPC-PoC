package publishers

import (
	"fmt"
	"github.com/juricaKenda/gRPC-PoC/topicbroker/infrastructure/pubsub"
	"time"
)

type timepub struct {
	subscribers []pubsub.Subscriber
	tag         string
}

func NewTimePublisher() pubsub.Publisher {
	return &timepub{
		subscribers: make([]pubsub.Subscriber, 0),
		tag:         "time_publisher",
	}
}

func (t *timepub) Start() {
	fmt.Println("time_publisher starting..")
	go t.run()
}

func (t *timepub) NotifyAll() {
	message := time.Now().String()

	for _, sub := range t.subscribers {
		sub.Receive(message)
	}
}

func (t *timepub) Subscribe(subscriber pubsub.Subscriber) {
	t.subscribers = append(t.subscribers, subscriber)
}

func (t *timepub) run() {
	for {
		t.NotifyAll()
		time.Sleep(2 * time.Second)

	}
}
