package appctx

import (
	"github.com/hieuus/food-delivery/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appCtx struct {
	db             *gorm.DB
	uploadprovider uploadprovider.UploadProvider
	secretKey      string
}

func NewAppContext(db *gorm.DB, uploadprovider uploadprovider.UploadProvider, secretKey string) *appCtx {
	return &appCtx{
		db:             db,
		uploadprovider: uploadprovider,
		secretKey:      secretKey,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB                 { return ctx.db }
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider { return ctx.uploadprovider }
func (ctx *appCtx) SecretKey() string                             { return ctx.secretKey }
