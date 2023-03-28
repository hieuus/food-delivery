package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	restaurantbiz "github.com/hieuus/food-delivery/module/restaurant/biz"
	restaurantmodel "github.com/hieuus/food-delivery/module/restaurant/model"
	restaurantstorage "github.com/hieuus/food-delivery/module/restaurant/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(context.Request.Context(), &data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
