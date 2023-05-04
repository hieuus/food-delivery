package restaurantlikebiz

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	restaurantlikemodel "github.com/hieuus/food-delivery/module/restaurantlike/model"
	"github.com/hieuus/food-delivery/pubsub"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

//type IncreaseLikedCountRestaurantStore interface {
//	IncreaseLikeCount(ctx context.Context, id int) error
//}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	//increase IncreaseLikedCountRestaurantStore
	ps pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	//increase IncreaseLikedCountRestaurantStore,
	ps pubsub.Pubsub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store,
		//increase,
		ps}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	//Send message
	if err := biz.ps.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	//Side Effect
	//j := asyncjob.NewJob(func(ctx context.Context) error {
	//	return biz.increase.IncreaseLikeCount(ctx, data.RestaurantId)
	//})
	//
	//if err := asyncjob.NewGroup(true, *j).Run(ctx); err != nil {
	//	log.Println(err)
	//}

	return nil
}
