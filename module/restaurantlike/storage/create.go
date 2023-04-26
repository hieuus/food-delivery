package restaurantlikestorage

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	restaurantlikemodel "github.com/hieuus/food-delivery/module/restaurantlike/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
