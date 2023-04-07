package uploadbusiness

import (
	"bytes"
	"context"
	"fmt"
	"github.com/hieuus/food-delivery/common"
	"github.com/hieuus/food-delivery/component/uploadprovider"
	"github.com/hieuus/food-delivery/module/upload/uploadmodel"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type uploadBiz struct {
	provider uploadprovider.UploadProvider
}

func NewUploadBiz(provider uploadprovider.UploadProvider) *uploadBiz {
	return &uploadBiz{provider: provider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, filename string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDemension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(filename)
	filename = fmt.Sprint("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, filename))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}
	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}

func getImageDemension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig((reader))

	if err != nil {
		log.Println("error", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
