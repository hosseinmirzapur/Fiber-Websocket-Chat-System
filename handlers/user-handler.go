package handlers

import "github.com/hosseinmirzapur/second_test_project/models"

func FindUserName(key string, users []models.User) string {
	for _, item := range users {
		if item.ID == key {
			return item.UserName
		}
	}
	return ""
}

var Users []models.User
