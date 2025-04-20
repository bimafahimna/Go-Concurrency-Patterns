package subscriber

import (
	"fmt"
	"pubsub/message_broker"
)

type Subscriber struct {
	id            int
	messageBroker message_broker.MessageBroker
}

func NewSubscriber(id int, msgBroker message_broker.MessageBroker) *Subscriber {
	return &Subscriber{
		id:            id,
		messageBroker: msgBroker,
	}
}

func (s *Subscriber) Subscribe(topic string) {
	ch := s.messageBroker.Subscribe(s.id, topic)
	go s.listen(ch)
}

func (s *Subscriber) listen(topic <-chan string) {
	for message := range topic {
		fmt.Printf("[Subs %d] %s\n", s.id, message)
		s.messageBroker.WaitGroupDone()
	}
}

func (s *Subscriber) Unsubscribe(topic string) {
	s.messageBroker.Unsubscribe(s.id, topic)
}
