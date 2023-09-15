package request

type ProductDescriptionRequest struct {
	Brand       string   `json:"brand"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	Name        string   `json:"name"`
}
