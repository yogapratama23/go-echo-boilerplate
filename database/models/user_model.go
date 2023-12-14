package models

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
