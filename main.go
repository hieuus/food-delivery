package main

import (
	"github.com/gin-gonic/gin"
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

	log.Println(db)

	//Simple CRUD
	//1. Create
	/*
		newRestaurant := Restaurant{Name: "Tani", Addr: "09 Pham Van hai"}

		if err := db.Create(&newRestaurant).Error; err != nil {
			log.Println(err)
		}
	*/

	//2. Read -> Where
	/*
		var myRestaurant Restaurant

		if err := db.Where("id = ?", 1).Find(&myRestaurant).Error; err != nil {
			log.Println(err)
		}

			log.Println(myRestaurant)
	*/

	//3. Update
	/*
		myRestaurant.Name = "Remove TaniPlus"
		if err := db.Where("id = ?", 3).Updates(&myRestaurant).Error; err != nil {
			log.Println(err)
		}

		log.Println(myRestaurant)
	*/

	//4. Delete
	/*
		if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 3).Delete(nil).Error; err != nil {
				log.Println(err)
			}
	*/

	//REST API
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")

	//1. Create new restaurant
	restaurants.POST("/", ginrestaurant.CreateRestaurant(db))

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

	//3. Get all and paging
	restaurants.GET("", func(context *gin.Context) {
		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}

		var pagingData Paging

		var data []Restaurant

		if err := context.ShouldBind(&pagingData); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}

		if pagingData.Limit <= 0 {
			pagingData.Limit = 5
		}

		if err := db.Order("id desc").
			Offset((pagingData.Page - 1) * pagingData.Limit).
			Limit(pagingData.Limit).
			Find(&data).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		context.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

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
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(db))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
