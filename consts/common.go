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
)
