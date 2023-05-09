package subscriber

import (
	"context"
	"github.com/hieuus/food-delivery/component/appctx"
	restaurantstorage "github.com/hieuus/food-delivery/module/restaurant/storage"
	"github.com/hieuus/food-delivery/pubsub"
)

//Pubsub
//func DecreaseLikeCountAfterUserDislikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubsub().Subcribe(ctx, common.TopicUserDislikeRestaurant)
//
//	store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
//
//	go func() {
//		defer common.AppRecover()
//		for {
//			msg := <-c
//			likeData := msg.Data().(HasRestaurantId)
//			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
//		}
//	}()
//}

// Engine: pubsub + asyncJob
func DecreaseLikeCountAfterUserDislikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		"Decrease like count after user dislikes restaurant",
		func(ctx context.Context, msg *pubsub.Message) error {
			store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := msg.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
