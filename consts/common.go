package consts

type XRequestID string

const (
	// Custom HTTP headers
	XRequestIDKey  = "X-Request-Id"
	XRequestIDCode = "xrId"
	XApiKey        = "x-api-key"

	// env
	DEVELOPMENT = "development"
	PRODUCTION  = "production"

	// Auth keys
	AuthUserIDKey = "id"

	// Request tag
	TagUserLoginRequest          = "userLoginRequest"
	TagProgressPercentageRequest = "progressPercentageRequest"
	TagAddUserImageRequest       = "addUserImageRequest"
	TagProductDescriptionRequest = "productDescriptionRequest"
	TagFileUploadRequest         = "fileUploadRequest"
)

// File Processing
func GetFileProcHeaders() []string {
	return []string{
		"SKU",
		"Name",
		"Stock Limit",
		"Version",
		"Retail Price",
		"Description",
		"How To Use",
		"Ingredients",
		"Product Images",
		"Height",
		"Width",
		"Length",
		"Weight",
	}
}
