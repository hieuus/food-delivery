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

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := context.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter

		if err := context.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.Status = []int{1}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(context.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		context.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
