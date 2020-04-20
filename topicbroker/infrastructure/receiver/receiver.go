package receiver

type Receiver struct {
	messages chan string
}

func NewMessageReceiver() *Receiver {
	return &Receiver{
		messages: make(chan string),
	}
}
func (r *Receiver) Receive(message string) {
	r.messages <- message
}

func (r *Receiver) Next() chan string {
	return r.messages
}
