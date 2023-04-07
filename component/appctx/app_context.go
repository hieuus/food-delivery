package appctx

import (
	"github.com/hieuus/food-delivery/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db             *gorm.DB
	uploadprovider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadprovider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{
		db:             db,
		uploadprovider: uploadprovider,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB                 { return ctx.db }
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider { return ctx.uploadprovider }
