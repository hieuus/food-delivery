package main

import (
	"context"
	"github.com/hieuus/food-delivery/pubsub"
	"github.com/hieuus/food-delivery/pubsub/localpubsub"
	"log"
	"time"
)

func main() {
	localPb := localpubsub.NewPubsub()
	var topic pubsub.Topic = "OrderCreated"

	sub1, close1 := localPb.Subcribe(context.Background(), topic)
	sub2, _ := localPb.Subcribe(context.Background(), topic)

	localPb.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPb.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func() {
		for {
			log.Println("Sub1:", (<-sub1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		for {
			log.Println("Sub2:", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)

	close1()
	localPb.Publish(context.Background(), topic, pubsub.NewMessage(3))
	time.Sleep(time.Second * 3)
}
