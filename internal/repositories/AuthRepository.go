package repositories

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/db"
	"fmt"
	"log"
)

func FindByUsername(login string) (*models.AuthUser, error) {
	if db.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return nil, fmt.Errorf("подключение к базе данных не инициализировано")
	}

	var user models.AuthUser
	query := db.DB.QueryRow("SELECT id, email, phone_number, password FROM auth_user WHERE email = $1 OR phone_number = $2", login, login)
	if err := query.Scan(&user.ID, &user.Email, &user.PhoneNumber, &user.Password); err != nil {
		log.Printf("Ошибка при сканировании данных пользователя: %v", err)
		return nil, err
	}
	return &user, nil
}
