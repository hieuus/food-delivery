package restaurantbiz

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	restaurantmodel "github.com/hieuus/food-delivery/module/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.repo.ListRestaurant(context, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
