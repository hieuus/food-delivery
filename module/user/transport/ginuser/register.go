package ginuser

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/component/appctx"
	"github.com/hieuus/food-delivery/component/hasher"
	userbiz "github.com/hieuus/food-delivery/module/user/biz"
	usermodel "github.com/hieuus/food-delivery/module/user/model"
	userstorage "github.com/hieuus/food-delivery/module/user/storage"
	"net/http"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterStorage(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
