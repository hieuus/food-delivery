package tokenprovider

import (
	"errors"
	"github.com/hieuus/food-delivery/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

type Token struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	Expiry    int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found", "ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding token"),
		"error encoding token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token provider"),
		"invalid token provider",
		"ErrInvalidToken",
	)
)
