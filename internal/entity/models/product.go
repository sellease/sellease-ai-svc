package models

type ProductDescription struct {
	ID         string `json:"id" gorm:"primaryKey"`
	PuzzleName string `json:"puzzle_name"`
}

type ProductDescriptionRequestData struct {
	Brand       string   `json:"brand"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	MaxTokens   int      `json:"max_tokens"`
	Model       string   `json:"model"`
	N           int      `json:"n"`
	Name        string   `json:"name"`
	SourceLang  string   `json:"source_lang"`
	TargetLang  string   `json:"target_lang"`
	Temperature float64  `json:"temperature"`
}
