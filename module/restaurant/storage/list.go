package restaurantstorage

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	restaurantmodel "github.com/hieuus/food-delivery/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())
	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}
		if len(filter.Status) > 0 {
			db.Where("status in (?)", filter.Status)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	limit := paging.Limit

	if err := db.Offset(offset).Limit(limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
