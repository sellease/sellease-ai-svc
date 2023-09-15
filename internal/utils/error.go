package utils

import "errors"

var (
	// request
	ErrInvalidRequest   = errors.New("unable to process request")
	ErrInvalidParameter = errors.New("invalid parameters")

	// user
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidUser  = errors.New("invalid user")

	// database
	ErrDBFailedToFetchData  = errors.New("error while reading data from db")
	ErrDBFailedToUpdateData = errors.New("error while updating data to db")
	ErrDBFailedToCreateData = errors.New("error while creating data to db")
	ErrDBFailedToDeleteData = errors.New("error while deleting data from db")

	// puzzle
	ErrPuzzleNotStarted = errors.New("puzzle has not been started")

	// auth
	ErrMissingAuthorization = errors.New("missing Authorization header or cookie")
	ErrInActiveUser         = errors.New("user is inactive")
)
