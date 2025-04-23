package message_broker

import (
	"sync"
)

type MessageBroker interface {
	Publish(topic, message string, useConcurrent bool)
	Subscribe(subsId int, topic string) <-chan string
	Unsubscribe(subsId int, topic string)
	WaitGroupAdd(i int)
	WaitGroupDone()
	WaitGroupWait()
}

type messageBroker struct {
	mu   *sync.RWMutex
	wg   *sync.WaitGroup
	subs map[string]map[int]chan string
}

func NewMessagebroker() MessageBroker {
	return &messageBroker{
		mu:   &sync.RWMutex{},
		wg:   &sync.WaitGroup{},
		subs: make(map[string]map[int]chan string),
	}
}

func (m *messageBroker) Publish(topic, message string, useConcurrent bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.WaitGroupAdd(len(m.subs[topic]))
	broadcast := func() {
		for _, ch := range m.subs[topic] {
			ch <- message
		}
	}

	if useConcurrent {
		go broadcast()
	} else {
		broadcast()
	}

}

func (m *messageBroker) Subscribe(subsId int, topic string) <-chan string {
	ch := make(chan string)

	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.subs[topic]
	if !ok {
		m.subs[topic] = make(map[int]chan string)
	}
	m.subs[topic][subsId] = ch
	return ch
}

func (m *messageBroker) Unsubscribe(subsId int, topic string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.subs[topic], subsId)
}

func (m *messageBroker) WaitGroupAdd(i int) {
	m.wg.Add(i)
}
func (m *messageBroker) WaitGroupDone() {
	m.wg.Done()
}
func (m *messageBroker) WaitGroupWait() {
	m.wg.Wait()
}
