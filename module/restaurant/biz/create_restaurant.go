package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "github.com/hieuus/food-delivery/module/restaurant/model"
)

type CreateRestaurantStore interface {
	CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("name cannot be empty")
	}

	if err := biz.store.CreateRestaurant(context, data); err != nil {
		return err
	}
	return nil
}
