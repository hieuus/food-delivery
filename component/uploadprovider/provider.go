package uploadprovider

import (
	"context"
	"github.com/hieuus/food-delivery/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
