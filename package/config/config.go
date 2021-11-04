package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type cfg struct {
	ApiKey string
}

var Cfg cfg

func Config() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error config!", err)
		return err
	}

	Cfg.ApiKey = os.Getenv("API_KEY")

	return nil
}
