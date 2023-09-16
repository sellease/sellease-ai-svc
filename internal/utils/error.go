package utils

import "errors"

var (
	// request
	ErrInvalidRequest   = errors.New("unable to process request")
	ErrInvalidParameter = errors.New("invalid parameters")

	// file proccesing
	ErrDescriptionTooShort       = errors.New("description too short")
	ErrProductImagesInsufficient = errors.New("atleast four product images must be provided")
	ErrAddingProductListing      = errors.New("error adding product listing")

	// translation
	ErrTranslation = errors.New("unable to translate text")
)
