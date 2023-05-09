package restaurantbiz

import (
	"context"
	"errors"
	"github.com/hieuus/food-delivery/common"
	restaurantmodel "github.com/hieuus/food-delivery/module/restaurant/model"
	"testing"
)

type mockCreateStore struct{}

func (mockCreateStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "Name" {
		return common.ErrDB(errors.New("something went wrong"))
	}

	data.Id = 200
	return nil
}

func TestNewCreateRestaurantBiz(t *testing.T) {
	biz := NewCreateRestaurantBiz(mockCreateStore{})

	dataTest := restaurantmodel.RestaurantCreate{Name: ""}
	err := biz.CreateRestaurant(context.Background(), &dataTest)

	if err == nil || err.Error() != "invalid request" {
		t.Errorf("Failed")
	}

	dataTest = restaurantmodel.RestaurantCreate{Name: "Name"}
	err = biz.CreateRestaurant(context.Background(), &dataTest)

	if err == nil {
		t.Errorf("Failed")
	}

	dataTest = restaurantmodel.RestaurantCreate{Name: "Right Name"}
	err = biz.CreateRestaurant(context.Background(), &dataTest)

	if err != nil {
		t.Errorf("Failed")
	}

}
