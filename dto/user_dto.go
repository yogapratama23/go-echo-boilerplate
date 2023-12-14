package dto

type FindManyUserFilter struct {
	IDs      []int32 `query:"ids"`
	Username string  `query:"username"`
	Fullname string  `query:"fullname"`
}

type GetUsersPayload struct {
	FindManyUserFilter
	PaginationInput
}
