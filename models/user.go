package models

type User struct {
	ID       string `json:"id"`
	UserName string `json:"username" validation:"required,min:1"`
}
