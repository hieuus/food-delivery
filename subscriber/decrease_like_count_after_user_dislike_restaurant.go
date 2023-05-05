package subscriber

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/component/appctx"
	restaurantstorage "github.com/hieuus/food-delivery/module/restaurant/storage"
)

func DecreaseLikeCountAfterUserDislikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subcribe(ctx, common.TopicUserLikeRestaurant)

	store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}
