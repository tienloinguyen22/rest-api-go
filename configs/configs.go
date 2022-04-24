package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Configs struct {
	ADDRESS string
	FIREBASE_CREDENTIALS_FILE string
	DB_URI string
	DB_STATEMENT_TIMEOUT time.Duration
}

func InitializeConfigs() *Configs {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
		os.Exit(1)
	}

	DB_STATEMENT_TIMEOUT := 10 * time.Second
	if duration, err := time.ParseDuration(os.Getenv("DB_STATEMENT_TIMEOUT")); err != nil {
		DB_STATEMENT_TIMEOUT = duration
	}

	configs := &Configs{
		ADDRESS: os.Getenv("ADDRESS"),
		FIREBASE_CREDENTIALS_FILE: os.Getenv("FIREBASE_CREDENTIALS_FILE"),
		DB_URI: os.Getenv("DB_URI"),
		DB_STATEMENT_TIMEOUT: DB_STATEMENT_TIMEOUT,
	}
	return configs
}

 