package userstorage

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	usermodel "github.com/hieuus/food-delivery/module/user/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, condition map[string]interface{}, moreInfor ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfor {
		db = db.Preload(moreInfor[i])
	}

	var user usermodel.User

	if err := db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
