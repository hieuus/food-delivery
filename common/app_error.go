package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, log, msg, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootError()
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"something went wrong in the server", err.Error(), "ErrInternal")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}
func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("CannotGet%s", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Already exists %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrEntityExisted%s", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}
