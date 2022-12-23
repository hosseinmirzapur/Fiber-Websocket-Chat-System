package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	envError := godotenv.Load()
	if envError != nil {
		log.Println(envError.Error())
	}
}
