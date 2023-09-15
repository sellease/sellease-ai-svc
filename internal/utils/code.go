package utils

const (
	InternalServerErrCode    = 500
	DatabaseErrCode          = 100
	ValidationErrCode        = 101
	InvalidParamsErrCode     = 102
	InvalidCredentialErrCode = 103
	AccessTokenErrCode       = 104
	PuzzleNotStartedErrCode  = 105
)

var ErrorCodeErrorMessage = map[int]string{
	InternalServerErrCode:    "ERR_INTERNAL_SERVER_ERROR",
	DatabaseErrCode:          "ERR_DB",
	ValidationErrCode:        "ERR_VALIDATION",
	InvalidParamsErrCode:     "ERR_INVALID_PARAMS",
	InvalidCredentialErrCode: "ERR_INVALID_CREDENTIAL",
	AccessTokenErrCode:       "ERR_GENERATING_ACCESS_TOKEN",
	PuzzleNotStartedErrCode:  "ERR_PUZZLE_NOT_STARTED",
}
