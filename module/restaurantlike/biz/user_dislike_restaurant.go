package restaurantlikebiz

import (
	"context"
	restaurantlikemodel "github.com/hieuus/food-delivery/module/restaurantlike/model"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, userId int, restaurantId int) error
}

type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(ctx context.Context, userId int, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	return nil
}
