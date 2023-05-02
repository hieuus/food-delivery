package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	id        string
	channle   Topic
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()

	return &Message{
		id:        fmt.Sprintf("%d", now.UnixNano()),
		data:      data,
		createdAt: now,
	}
}

func (evt *Message) String() string {
	return fmt.Sprintf("Message %s", evt.channle)
}

func (evt *Message) Channel() Topic {
	return evt.channle
}

func (evt *Message) SetChannel(channel Topic) {
	evt.channle = channel
}

func (evt *Message) Data() interface{} {
	return evt.data
}
