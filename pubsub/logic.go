package pubsub

import (
	"fmt"
	"sync"
	"time"
)

type Pubsub struct {
	mu   sync.RWMutex
	subs map[string][]chan string
}

//constructor
func NewPubsub() *Pubsub {
	ps := &Pubsub{}
	ps.subs = make(map[string][]chan string)
	return ps
}

func (ps *Pubsub) AddSubscription(topic string, ch chan string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.subs[topic] = append(ps.subs[topic], ch)
}

// func (ps *Pubsub) DeleteSubscription(ch chan string) {
// 	close(ch)
// }

func (ps *Pubsub) Publish(topic string, msg string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for _, ch := range ps.subs[topic] {
		ch <- msg
		ack := <-ch
		fmt.Printf("Get acknowleged from @%s\n", ack)

	}
}

func (ps *Pubsub) CreateTopic(topic string) chan string {
	ch := make(chan string, 1)
	ps.AddSubscription(topic, ch)
	return ch
}

func (ps *Pubsub) Ack(name string, ch chan string) {
	ch <- name
}

func (ps *Pubsub) DeleteTopic(topic string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	delete(ps.subs, topic)
}

func (ps *Pubsub) Subscribe(name string, ch chan string) {
	for i := range ch {
		fmt.Printf("[%s] got %s\n", name, i)
		ps.Ack(name, ch)

	}
}

func (ps *Pubsub) Pub(topic string, msg string) {
	fmt.Printf("Publishing @%s: %s\n", topic, msg)
	ps.Publish(topic, msg)
	time.Sleep(1 * time.Millisecond)
}
