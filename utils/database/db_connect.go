package database

import (
	"STDE_proj/configs"
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

// Connect устанавливает соединение с базой данных PostgreSQL и инициализирует глобальную переменную DB.
// Строка соединения формируется с использованием параметров, заданных в конфигурации AppConfig.
// Возвращает ошибку, если не удалось подключиться к базе данных или проверить соединение.
func Connect() error {
	var err error

	// Формируем строку соединения с базой данных
	con := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		configs.AppConfig.Database.User, configs.AppConfig.Database.Password,
		configs.AppConfig.Database.Name, configs.AppConfig.Database.Host,
		configs.AppConfig.Database.Port)

	// Открываем соединение с базой данных
	DB, err = sql.Open("postgres", con)
	if err != nil {
		return fmt.Errorf("Ошибка подключения к базе данных: %w", err)
	}

	// Проверяем соединение с базой данных
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("Ошибка при проверке подключения к базе данных: %w", err)
	}

	log.Println("Подключение к базе данных успешно!")
	return nil
}
