package restaurantlikebiz

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	restaurantlikemodel "github.com/hieuus/food-delivery/module/restaurantlike/model"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncreaseLikedCountRestaurantStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	increase IncreaseLikedCountRestaurantStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, increase IncreaseLikedCountRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store, increase}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()
		if err := biz.increase.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
