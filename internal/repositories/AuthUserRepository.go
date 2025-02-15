package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/db"
	"errors"
	"fmt"
	"log"
)

func PostAuthUser(data models.AuthUser) error {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return errors.New("подключение к базе данных не инициализировано")
	}

	query := "INSERT INTO auth_user (password, email, phone_number) VALUES ($1, $2, $3)"
	result, err := db.DB.Exec(query, data.Password, data.Email, data.PhoneNumber)
	if err != nil {
		log.Printf("Ошибка при добавлении: %v", err)
		return fmt.Errorf("ошибка при добавлении: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества добавленных строк: %v", err)
		return fmt.Errorf("ошибка при проверке количества добавленных строк: %v", err)
	}
	if rowsAffected == 0 {
		log.Println("Ошибка: пользователь не был добавлен")
		return errors.New("пользователь не был добавлен")
	}

	return nil
}
