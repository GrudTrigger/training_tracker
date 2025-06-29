package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Configs struct {
	Port string
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
		Db{Dsn: os.Getenv("DSN")},
	}
}
