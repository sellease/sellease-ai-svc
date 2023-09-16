package request

type ProductDescriptionRequest struct {
	// Brand       string   `json:"brand,omitempty"`
	// Category    string   `json:"category,omitempty"`
	// Description string   `json:"description,omitempty"`
	// Keywords    []string `json:"keywords"`
	Name string `json:"name"`
}

type TranslationRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
	HowToUse    string `json:"how_to_use"`
	Language    string `json:"language"`
}
