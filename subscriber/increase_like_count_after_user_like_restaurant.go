package subscriber

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/component/appctx"
	restaurantstorage "github.com/hieuus/food-delivery/module/restaurant/storage"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	//GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subcribe(ctx, common.TopicUserLikeRestaurant)

	store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}
