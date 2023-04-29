package restaurantlikebiz

import (
	"context"
	"github.com/hieuus/food-delivery/component/asyncjob"
	restaurantlikemodel "github.com/hieuus/food-delivery/module/restaurantlike/model"
	"log"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, userId int, restaurantId int) error
}

type DecreaseLikedCountRestaurantStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userDislikeRestaurantBiz struct {
	store         UserDislikeRestaurantStore
	decreaseStore DecreaseLikedCountRestaurantStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore, decStore DecreaseLikedCountRestaurantStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store, decreaseStore: decStore}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(ctx context.Context, userId int, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	//Side Effect
	j := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.decreaseStore.DecreaseLikeCount(ctx, restaurantId)
	})

	if err := asyncjob.NewGroup(true, *j).Run(ctx); err != nil {
		log.Println(err)
	}

	//go func() {
	//	defer common.AppRecover()
	//	if err := biz.decreaseStore.DecreaseLikeCount(ctx, userId); err != nil {
	//		log.Println(err)
	//	}
	//}()

	return nil
}
