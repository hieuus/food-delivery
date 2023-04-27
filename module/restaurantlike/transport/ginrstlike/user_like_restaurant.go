package ginrstlike

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/component/appctx"
	restaurantlikebiz "github.com/hieuus/food-delivery/module/restaurantlike/biz"
	restaurantlikemodel "github.com/hieuus/food-delivery/module/restaurantlike/model"
	restaurantlikestorage "github.com/hieuus/food-delivery/module/restaurantlike/storage"
	"net/http"
)

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
