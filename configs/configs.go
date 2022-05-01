package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	ADDRESS string
	FIREBASE_CREDENTIALS_FILE string
	DB_URI string
	REDIS_URI string
}

func InitializeConfigs() *Configs {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
		os.Exit(1)
	}

	configs := &Configs{
		ADDRESS: os.Getenv("ADDRESS"),
		FIREBASE_CREDENTIALS_FILE: os.Getenv("FIREBASE_CREDENTIALS_FILE"),
		DB_URI: os.Getenv("DB_URI"),
		REDIS_URI: os.Getenv("REDIS_URI"),
	}
	return configs
}

 