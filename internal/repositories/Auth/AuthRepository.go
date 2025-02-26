package Auth

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"STDE_proj/utils/smtp_sender"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

func FindByUsername(data models.AuthUser) (*models.AuthUser, error) {

	switch data.TypeLogin {
	case "email":
		query := database.DB.QueryRow(`
        SELECT id, email, password 
        FROM auth_user 
        WHERE email = $1`, data.Login)
		err := query.Scan(&data.ID, &data.Login, &data.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				// Пользователь не найден
				log.Printf("Пользователь с логином %s не найден", data.Login)
				return nil, fmt.Errorf("пользователь не найден")
			}
			// Другие ошибки (например, проблемы с подключением или сканированием)
			log.Printf("Ошибка при сканировании данных пользователя: %v", err)
			return nil, fmt.Errorf("ошибка при поиске пользователя: %v", err)
		}
		return &data, nil
	case "phone_number":
		query := database.DB.QueryRow(`
        SELECT id, phone_number, password 
        FROM auth_user 
        WHERE phone_number = $1`, data.Login)
		err := query.Scan(&data.ID, &data.Login, &data.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				// Пользователь не найден
				log.Printf("Пользователь с логином %s не найден", data.Login)
				return nil, fmt.Errorf("пользователь не найден")
			}
			// Другие ошибки (например, проблемы с подключением или сканированием)
			log.Printf("Ошибка при сканировании данных пользователя: %v", err)
			return nil, fmt.Errorf("ошибка при поиске пользователя: %v", err)
		}
		return &data, nil
	}
	return &models.AuthUser{}, errors.New("некорректный логин пользователя")
}

// В будущем навзать CheckVerifyAccount
// Добавь возможность проверять аккаунт подтвержден или нет не только почтой, но и номером телефона
func CheckVerifyEmail(email string, idAuthUser int) error {

	query := database.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM auth_user WHERE email = $1 AND is_email_verify = true);", email)
	var user models.AuthUser

	if err := query.Scan(&user.IsEmail); err != nil {
		log.Printf("Ошибка: %v", err)
		return errors.New("ошибка при попытки принять значение с базы данных")
	}
	if user.IsEmail == false {
		return errors.New("почта не подтверждена")
	}
	verifyCode := smtp_sender.GenerateCode()
	_, err := database.DB.Exec("INSERT INTO verify_code (code, auth_user_id) VALUES ($1, $2);", verifyCode, idAuthUser)
	if err != nil {
		log.Printf("ошибка при создании проверочного кода %v", err)
		return errors.New("ошибка при создании проверочного кода")
	}

	err = smtp_sender.SendEmail(email, verifyCode)
	if err != nil {
		log.Fatalf("Ошибка при отправке письма: %v", err)
	}

	return nil
}

func CheckVerifyPhoneNumber(phoneNumber string) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	query := database.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM auth_user WHERE phone_number = $1 AND is_number_verify = true);", phoneNumber)
	var user models.AuthUser

	if err := query.Scan(&user.IsEmail); err != nil {
		log.Printf("Ошибка: %v", err)
		return errors.New("ошибка при попытки принять значение с базы данных")
	}
	if user.IsEmail == false {
		return errors.New("номер телефона не подтвержден")
	}
	return nil
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
