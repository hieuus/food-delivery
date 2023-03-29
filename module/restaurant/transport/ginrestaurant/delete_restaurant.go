package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/component/appctx"
	restaurantbiz "github.com/hieuus/food-delivery/module/restaurant/biz"
	restaurantstorage "github.com/hieuus/food-delivery/module/restaurant/storage"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()

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

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
