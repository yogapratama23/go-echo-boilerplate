package models

type User struct {
	BaseModel
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}
