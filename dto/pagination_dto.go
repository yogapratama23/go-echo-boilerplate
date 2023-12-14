package dto

type PaginationInput struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PaginationResponse struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	Total     int `json:"total"`
	Limit     int `json:"limit"`
}

type PaginationOptions struct {
	Limit  int
	Offset int
}
