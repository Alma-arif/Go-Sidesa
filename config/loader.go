package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error when load env: %s", err.Error())
	}

	return Config{
		Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_Port"),
		},
		Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
		Enkripsi{
			Key: os.Getenv("KEY"),
		},
		FileConf{
			ImageDerektory: os.Getenv("IMAGE_DEREKTORY"),
			FileDerektory:  os.Getenv("FILE_DEREKTORY"),
			FileType:       os.Getenv("FILE_TYPE"),
		},
	}

}
