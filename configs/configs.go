package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type configs struct {
	ADDRESS string
}

var Configs configs

func InitializeConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	Configs.ADDRESS = os.Getenv("ADDRESS")
}

 