package localpubsub

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/pubsub"
	"log"
	"sync"
)

type localPubsub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewPubsub() *localPubsub {
	pb := &localPubsub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (ps *localPubsub) Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error {
	data.SetChannel(topic)

	go func() {
		defer common.AppRecover()
		ps.messageQueue <- data
		log.Println("New event published", data.String(), "with data", data.Data())
	}()

	return nil
}

func (ps *localPubsub) Subcribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock()

	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubcribe")

		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()

					break
				}
			}
		}
	}
}

func (ps *localPubsub) run() error {
	log.Println("Pubsub started")

	go func() {
		defer common.AppRecover()

		for {
			msg := <-ps.messageQueue
			log.Println("Message dequeue", msg)

			if subs, ok := ps.mapChannel[msg.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						defer common.AppRecover()
						c <- msg
					}(subs[i])
				}
			}
		}

	}()

	return nil
}
