package response

type ProcessFileResponse struct {
	SuccessCount int64
	ErrorCount   int64
	Errors       []FileError
}

type FileError struct {
	SKU   string
	Error string
}
