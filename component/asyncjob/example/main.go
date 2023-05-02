package main

import (
	"context"
	"github.com/hieuus/food-delivery/component/asyncjob"
	"log"
	"time"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("I am job 1")
		return nil
		//return errors.New("Something went wrong in Job 1")
	})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("I am job 2")
		return nil
		//return errors.New("Something went wrong in Job 2")
	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("I am job 3")
		return nil
		//return errors.New("Something went wrong in Job 3")
	})

	group := asyncjob.NewGroup(true, *job1, *job2, *job3)

	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}
}
