package models

type ProductListing struct {
	SKU           string         `json:"sku"`
	Name          string         `json:"name"`
	StockLimit    int            `json:"stockLimit"`
	Version       int            `json:"version"`
	RetailPrice   float64        `json:"retailPrice"`
	Description   string         `json:"description"`
	HowToUse      string         `json:"howToUse"`
	Ingredients   string         `json:"ingredients"`
	ProductImages []ProductImage `json:"productImages"`
	Height        float64        `json:"height"`
	Width         float64        `json:"width"`
	Length        float64        `json:"length"`
	Weight        float64        `json:"weight"`
	BrandID       string         `json:"brandId"`
	CategoryID    string         `json:"categoryId"`
}

type ProductImage struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}
