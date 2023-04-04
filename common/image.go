package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string { return "image" }

func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to Unmarshal JSON value:", value))
	}

	var img Image

	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img

	return nil
}

func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

type Images []Image

func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to Unmarshal JSON value:", value))
	}

	var imgs []Image
	if err := json.Unmarshal(bytes, &imgs); err != nil {
		return err
	}

	*j = imgs
	return nil
}
func (j *Images) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
