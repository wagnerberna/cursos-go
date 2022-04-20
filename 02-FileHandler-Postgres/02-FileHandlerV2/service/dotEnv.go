package service

import (
	"os"

	"github.com/joho/godotenv"
)

func DotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}
