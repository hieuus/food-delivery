package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuus/food-delivery/component/appctx"
	"github.com/hieuus/food-delivery/component/uploadprovider"
	"github.com/hieuus/food-delivery/middleware"
	"github.com/hieuus/food-delivery/module/restaurant/transport/ginrestaurant"
	"github.com/hieuus/food-delivery/module/upload/uploadtransport/ginupload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	//Connect to DB
	dsn := os.Getenv("MYSQL_CONN_STRING")

	//aws s3
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := appctx.NewAppContext(db, s3Provider)

	//REST API
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.Upload(appCtx))

	restaurants := v1.Group("/restaurants")

	//1. Create new restaurant
	restaurants.POST("/", ginrestaurant.CreateRestaurant(appCtx))

	//2. GET By Id
	restaurants.GET("/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		var data Restaurant

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	//3. Get list restaurant with paging
	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

	//4. Update
	restaurants.PATCH("/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data RestaurantUpdate

		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	})

	//5 Delete
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
