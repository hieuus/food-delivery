package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuus/food-delivery/component/appctx"
	"github.com/hieuus/food-delivery/middleware"
	"github.com/hieuus/food-delivery/module/restaurant/transport/ginrestaurant"
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	appCtx := appctx.NewAppContext(db)

	//REST API
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
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
