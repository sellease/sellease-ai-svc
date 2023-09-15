package utils

import "errors"

var (
	// request
	ErrInvalidRequest   = errors.New("unable to process request")
	ErrInvalidParameter = errors.New("invalid parameters")

	// database
	ErrDBFailedToFetchData  = errors.New("error while reading data from db")
	ErrDBFailedToUpdateData = errors.New("error while updating data to db")
	ErrDBFailedToCreateData = errors.New("error while creating data to db")
	ErrDBFailedToDeleteData = errors.New("error while deleting data from db")

	// puzzle
	ErrPuzzleNotStarted = errors.New("puzzle has not been started")

	// file proccesing
	ErrDescriptionTooShort       = errors.New("description too short")
	ErrProductImagesInsufficient = errors.New("atleast four product images must be provided")
	ErrAddingProductListing      = errors.New("error adding product listing")
)
