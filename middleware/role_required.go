package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/component/appctx"
)

func RoleRequired(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		for _, item := range allowRoles {
			if item == u.GetUserRole() {
				c.Next()
				break
			}
		}

		panic(common.ErrNoPermission(errors.New("invalid role user")))
	}
}
