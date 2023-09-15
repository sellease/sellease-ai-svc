package response

type Output struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
	ID    string `json:"id"`
}

type ProductDescriptionData struct {
	Outputs          []Output `json:"outputs"`
	RemainingCredits float64  `json:"remaining_credits"`
}

type ProductDescriptionResponse struct {
	Status string                 `json:"status"`
	Data   ProductDescriptionData `json:"data"`
}
