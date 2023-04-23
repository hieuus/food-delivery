package common

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	GetUserRole() string
	GetUserEmail() string
}
