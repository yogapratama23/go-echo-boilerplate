package models

import "github.com/yogapratama23/go-echo-boilerplate/dto"

type User struct {
	BaseModel
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

type APIUser struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

type APIUserPaginate struct {
	Users              []*APIUser `json:"users"`
	PaginationResponse dto.PaginationResponse
}
