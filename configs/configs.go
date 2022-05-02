package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configs struct {
	ADDRESS string
	FIREBASE_CREDENTIALS_FILE string
	DB_URI string
	REDIS_URI string
	EMAIL_HOST string
	EMAIL_PORT int
	EMAIL_USERNAME string
	EMAIL_PASSWORD string
}

func InitializeConfigs() *Configs {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
		os.Exit(1)
	}

	EMAIL_PORT, err := strconv.ParseInt(os.Getenv("EMAIL_PORT"), 10, 64)
	if err != nil {
		EMAIL_PORT = 465 // Default value
	}

	configs := &Configs{
		ADDRESS: os.Getenv("ADDRESS"),
		FIREBASE_CREDENTIALS_FILE: os.Getenv("FIREBASE_CREDENTIALS_FILE"),
		DB_URI: os.Getenv("DB_URI"),
		REDIS_URI: os.Getenv("REDIS_URI"),
		EMAIL_HOST: os.Getenv("EMAIL_HOST"),
		EMAIL_PORT: int(EMAIL_PORT),
		EMAIL_USERNAME: os.Getenv("EMAIL_USERNAME"),
		EMAIL_PASSWORD: os.Getenv("EMAIL_PASSWORD"),
	}
	return configs
}

 