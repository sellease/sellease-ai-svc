package response

type Health struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
