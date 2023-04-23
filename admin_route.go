package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuus/food-delivery/component/appctx"
	"github.com/hieuus/food-delivery/middleware"
	"github.com/hieuus/food-delivery/module/user/transport/ginuser"
)

func setupAdminRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin",
		middleware.RequiredAuth(appCtx),
		middleware.RoleRequired(appCtx, "admin", "mod"),
	)

	{
		admin.GET("/profile", ginuser.Profile(appCtx))
	}
}
