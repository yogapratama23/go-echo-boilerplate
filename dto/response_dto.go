package dto

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponsePagination struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	PaginationResponse
}
