package common

type Pagination struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
	Skip      bool   `json:"skip"`
}

type PaginationWithFilters struct {
	Pagination
	Filters map[string]string `json:"filters,omitempty"`
}
