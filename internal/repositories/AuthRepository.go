package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func FindByUsername(login string) (*models.AuthUser, error) {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, fmt.Errorf("подключение к базе данных не инициализировано")
	}

	var user models.AuthUser
	query := database.DB.QueryRow("SELECT id, email, phone_number, password FROM auth_user WHERE email = $1 OR phone_number = $2", login, login)
	var phoneNumber *sql.NullString
	if err := query.Scan(&user.ID, &user.Email, &phoneNumber, &user.Password); err != nil {
		log.Printf("Ошибка при сканировании данных пользователя: %v", err)
		return nil, err
	}
	return &user, nil
}

func UpdateLastLogin(data *models.AuthUser) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	var currentTime = time.Now()

	query, err := database.DB.Exec("UPDATE auth_user SET last_login = $1 WHERE id = $2", currentTime, data.ID)
	if err != nil {
		log.Printf("Ошибка при обновлении время захода с id %d: %v", data.ID, err)
		return fmt.Errorf("ошибка при обновлении время захода: %v", err)
	}

	rowsAffected, err := query.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при проверке количества обновленных строк для id %d: %v", data.ID, err)
		return fmt.Errorf("ошибка при проверке количества обновленных строк: %v", err)
	}
	if rowsAffected == 0 {
		log.Printf("Ошибка: время захода с id %d не найдена", data.ID)
		return fmt.Errorf("время захода с id %d не найдена", data.ID)
	}

	return nil
}
