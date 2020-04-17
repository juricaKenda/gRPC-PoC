package pubsub

type Publisher interface {
	Start()
	NotifyAll()
	Subscribe(subscriber Subscriber)
}

type Subscriber interface {
	Receive(message, pubTag string)
}
