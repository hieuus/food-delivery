package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuus/food-delivery/component/appctx"
	"github.com/hieuus/food-delivery/middleware"
	"github.com/hieuus/food-delivery/module/restaurant/transport/ginrestaurant"
	"github.com/hieuus/food-delivery/module/upload/uploadtransport/ginupload"
	"github.com/hieuus/food-delivery/module/user/transport/ginuser"
	"net/http"
	"strconv"
)

func setupRoute(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/upload", ginupload.Upload(appCtx))

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/authenticate", ginuser.Login(appCtx))
	v1.GET("profile", middleware.RequiredAuth(appCtx), ginuser.Profile(appCtx))

	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appCtx))

	//1. Create new restaurant
	restaurants.POST("/", ginrestaurant.CreateRestaurant(appCtx))

	//2. GET By Id
	restaurants.GET("/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		var data Restaurant

		if err := appCtx.GetMainDBConnection().Where("id = ?", id).First(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	//3. Get list restaurant with paging
	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

	//4. Update
	restaurants.PATCH("/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data RestaurantUpdate

		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
		if err := appCtx.GetMainDBConnection().Where("id = ?", id).Updates(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	})

	//5 Delete
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
}
