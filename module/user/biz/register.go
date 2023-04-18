package userbiz

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	usermodel "github.com/hieuus/food-delivery/module/user/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfor ...string) (*usermodel.User, error)
	CreateUser(stc context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterStorage(registerBiz RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerBiz,
		hasher:          hasher,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
