package db

import (
	"STDE_proj/configs"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
)

var DB *sql.DB

func Connect() error {
	var err error

	con := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		configs.AppConfig.Database.User, configs.AppConfig.Database.Password,
		configs.AppConfig.Database.Name, configs.AppConfig.Database.Host,
		configs.AppConfig.Database.Port)

	DB, err = sql.Open("postgres", con)
	if err != nil {
		return fmt.Errorf("Ошибка подключения к базе данных: %w", err)
	}
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("Ошибка при проверке подключения к базе данных: %w", err)
	}

	log.Println("Подключение к базе данных успешно!")
	return nil
}
