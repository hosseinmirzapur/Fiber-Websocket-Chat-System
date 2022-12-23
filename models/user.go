package models

type User struct {
	ID       string `json:"id"`
	UserName string `json:"user_name" validation:"required,min:1"`
}
