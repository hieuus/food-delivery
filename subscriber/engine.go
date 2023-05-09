package subscriber

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/component/appctx"
	"github.com/hieuus/food-delivery/component/asyncjob"
	"github.com/hieuus/food-delivery/pubsub"
	"log"
)

type consumerJob struct {
	Title string
	Hdl   func(ctx context.Context, msg *pubsub.Message) error
}

type consumerEngine struct {
	appCtx appctx.AppContext
}

func NewEngine(appContext appctx.AppContext) *consumerEngine {
	return &consumerEngine{appContext}
}

func (engine *consumerEngine) Start() error {

	engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		IncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
		PushNotificationWhenUserLikeRestaurant(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicUserDislikeRestaurant,
		true,
		DecreaseLikeCountAfterUserDislikeRestaurant(engine.appCtx),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubsub().Subcribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("Setup subscriber for ", item.Title)
	}

	getJobHandler := func(job *consumerJob, msg *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running for ", job.Title, ". Value: ", msg.Data())
			return job.Hdl(ctx, msg)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}

	}()

	return nil
}
