package publisher

import (
	"errors"
	"fmt"
	"log"
	"pubsub/message_broker"
)

type Publisher struct {
	id            int
	messageBroker message_broker.MessageBroker
	useConcurrent bool
}

func NewPublisher(id int, msgBroker message_broker.MessageBroker, useConcurrent bool) *Publisher {
	return &Publisher{
		id:            id,
		messageBroker: msgBroker,
		useConcurrent: useConcurrent,
	}
}

func (p *Publisher) Publish(topic, message string) error {
	valid := p.validateMessage(message)
	if !valid {
		err := errors.New("invalid mesage")
		log.Println(err)
		return err
	}

	p.messageBroker.Publish(topic, fmt.Sprintf("From Pubs %d; topic: \"%s\" message: %s", p.id, topic, message), p.useConcurrent)
	return nil
}

func (p *Publisher) validateMessage(message string) bool {
	return len(message) >= 5 && len(message) <= 30
}
