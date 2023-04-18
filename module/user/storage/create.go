package userstorage

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	usermodel "github.com/hieuus/food-delivery/module/user/model"
)

func (s *sqlStore) CreateUser(stc context.Context, data *usermodel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
