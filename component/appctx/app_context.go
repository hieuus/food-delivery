package appctx

import (
	"github.com/hieuus/food-delivery/component/uploadprovider"
	"github.com/hieuus/food-delivery/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubsub() pubsub.Pubsub
}

type appCtx struct {
	db             *gorm.DB
	uploadprovider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, uploadprovider uploadprovider.UploadProvider, secretKey string, ps pubsub.Pubsub) *appCtx {
	return &appCtx{
		db:             db,
		uploadprovider: uploadprovider,
		secretKey:      secretKey,
		ps:             ps,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB                 { return ctx.db }
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider { return ctx.uploadprovider }
func (ctx *appCtx) SecretKey() string                             { return ctx.secretKey }
func (ctx *appCtx) GetPubsub() pubsub.Pubsub                      { return ctx.ps }
