package db

import (
	"log"
	"pingPong/internal/config"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	cfg := config.LoadConfig()
	dsn := cfg.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}

	log.Println("Успешное подключение к базе данных")

	return db
}

func LoadEnv() { // первоначальный вызов .env, чтобы Go мог читать переменные database
	err := godotenv.Load() // обработка ошибки, на тот случай, если не получается считывать данные, Go - будет использовать системный env`шник

	if err != nil {
		log.Println(".env не найден, используется системный .env")
	}
}
