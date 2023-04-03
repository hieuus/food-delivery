package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/component/appctx"
	restaurantbiz "github.com/hieuus/food-delivery/module/restaurant/biz"
	restaurantmodel "github.com/hieuus/food-delivery/module/restaurant/model"
	restaurantstorage "github.com/hieuus/food-delivery/module/restaurant/storage"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data restaurantmodel.RestaurantCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
