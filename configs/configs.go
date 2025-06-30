package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	Port string
	Secret string
	Db
}

type Db struct {
	Dsn string
}

func LoadConfigs() *Configs {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузку .env файла")
	}

	return &Configs{
		os.Getenv("PORT"),
		os.Getenv("SECRET"),
		Db{Dsn: os.Getenv("DSN")},
	}
}
