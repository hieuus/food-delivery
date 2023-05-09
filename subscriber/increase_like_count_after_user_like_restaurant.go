package subscriber

import (
	"context"
	"github.com/hieuus/food-delivery/component/appctx"
	restaurantstorage "github.com/hieuus/food-delivery/module/restaurant/storage"
	"github.com/hieuus/food-delivery/pubsub"
	"log"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	//GetUserId() int
}

//Pubsub
//func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubsub().Subcribe(ctx, common.TopicUserLikeRestaurant)
//
//	store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
//
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c
//			likeData := msg.Data().(HasRestaurantId)
//			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
//		}
//	}()
//}

// Engine: pubsub + asyncJob
func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		"Increase like count after user likes restaurant",
		func(ctx context.Context, msg *pubsub.Message) error {
			store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := msg.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

// Another consumer
func PushNotificationWhenUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		"Increase like count after user likes restaurant",
		func(ctx context.Context, msg *pubsub.Message) error {
			likeData := msg.Data().(HasRestaurantId)
			log.Println("Push notification when user likes restaurant id: ", likeData.GetRestaurantId())
			return nil
		},
	}
}
