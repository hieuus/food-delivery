package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	restaurantbiz "github.com/hieuus/food-delivery/module/restaurant/biz"
	restaurantstorage "github.com/hieuus/food-delivery/module/restaurant/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(context.Request.Context(), id); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"data": 1,
		})
	}
}
