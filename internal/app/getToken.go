package app

import (
	"github.com/joho/godotenv"
	"os"
)

func GetToken() string {
	err := godotenv.Load(".env")
	HandleError(err)

	return os.Getenv("TOKEN")
}
