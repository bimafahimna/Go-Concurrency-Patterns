package main

import (
	broker "pubsub/message_broker"
	pub "pubsub/publisher"
	sub "pubsub/subscriber"
)

func main() {
	messageBroker := broker.NewMessagebroker()

	publisher1 := pub.NewPublisher(1, messageBroker, false)
	publisher2 := pub.NewPublisher(2, messageBroker, true)

	subscriber1 := sub.NewSubscriber(1, messageBroker)
	subscriber4 := sub.NewSubscriber(4, messageBroker)

	subscriber1.Subscribe("foo")
	subscriber1.Subscribe("bar")
	subscriber1.Subscribe("foobar")
	subscriber4.Subscribe("foo")

	publisher1.Publish("bar", "Delicious Meal")
	publisher1.Publish("foobar", "dlrow olleH")
	publisher1.Publish("foo", "hello world")

	publisher2.Publish("foobar", "Easy Concurrent")
	publisher2.Publish("buzz", "Buzz buzz buzz")

	subscriber1.Unsubscribe("foo")

	publisher2.Publish("foo", "For subscriber only")

	messageBroker.WaitGroupWait()
}
