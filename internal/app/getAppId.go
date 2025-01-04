package app

import (
	"github.com/joho/godotenv"
	"os"
)

func GetAppId() string {
	err := godotenv.Load(".env")
	HandleError(err)

	return os.Getenv("APP_ID")
}
