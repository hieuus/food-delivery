package subscriber

import (
	"context"
	"github.com/hieuus/food-delivery/component/appctx"
)

func Setup(appCtx appctx.AppContext, ctx context.Context) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, ctx)
	DecreaseLikeCountAfterUserDislikeRestaurant(appCtx, ctx)
}
