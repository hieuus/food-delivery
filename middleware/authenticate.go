package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/component/appctx"
	"github.com/hieuus/food-delivery/component/tokenprovider/jwt"
	userstorage "github.com/hieuus/food-delivery/module/user/storage"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprint("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFormHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

//1. Get token from header
//2. Validate token and parse to payload
//3. From the token payload, we use user_id to find from DB

func RequiredAuth(appCtx appctx.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFormHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSqlStore(db)

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
