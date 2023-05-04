package common

import "log"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

const (
	CurrentUser = "user"
)

const TopicUserLikeRestaurant = "UserLikeRestaurant"

type Requester interface {
	GetUserId() int
	GetUserRole() string
	GetUserEmail() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error", err)
	}
}
