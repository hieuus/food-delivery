package userbiz

import (
	"context"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/component/tokenprovider"
	usermodel "github.com/hieuus/food-delivery/module/user/model"
	"log"
)

type LoginStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfor ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

//Flow
//1. Find User, email
//2. Hash pwd form input then compare with pwd in db
//3. Provider: JWT token -> Access token and refresh token
//4. Return token(s)

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	//1. Find user
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	log.Println(data.Email)
	log.Println(data.Password)
	log.Println(err)
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	//2.Hash pwd and compare
	pwdHashed := business.hasher.Hash(data.Password + user.Salt)

	if pwdHashed != user.Password {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	//3. JWT Token
	payLoad := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payLoad, business.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
