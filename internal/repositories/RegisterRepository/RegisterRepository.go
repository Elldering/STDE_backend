package RegisterRepository

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"STDE_proj/utils/smtp_sender"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func Register(registerData models.Register) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}
	var AuthUserData models.AuthUser
	var phoneNumber sql.NullString
	query := database.DB.QueryRow("SELECT id, email, phone_number FROM auth_user WHERE email = $1 OR phone_number = $2", registerData.Login, registerData.Login)
	err := query.Scan(&AuthUserData.ID, &AuthUserData.Email, &phoneNumber)
	if err == nil {
		// Пользователь с указанным email или номером телефона уже существует
		log.Println("Пользователь с таким email или номером телефона уже существует")
		return fmt.Errorf("пользователь с таким email или номером телефона уже существует")
	}

	if !errors.Is(err, sql.ErrNoRows) {
		// Обработка ошибки (если она не связана с отсутствием пользователя в БД)
		log.Printf("Ошибка при сканировании данных пользователя: %v", err)
		return err
	}

	// SQL-запрос для вставки данных в таблицу auth_user
	execAuthUser := "INSERT INTO auth_user (email, password) VALUES ($1, $2) RETURNING id"
	err = database.DB.QueryRow(execAuthUser, registerData.Login, registerData.Password).Scan(&AuthUserData.ID)
	if err != nil {
		log.Printf("Ошибка при добавлении пользователя в базу данных: %v", err)
		return err
	}
	log.Println("Пользователь успешно зарегистрирован")

	verifyCode := smtp_sender.GenerateCode()

	err = smtp_sender.SendEmail(registerData.Login, verifyCode)
	if err != nil {
		log.Fatalf("Ошибка при отправке письма: %v", err)
	}

	// SQL-запрос для вставки данных в таблицу verify_code
	execVerifyCode := "INSERT INTO verify_code (code, auth_user_id) VALUES ($1, $2)"
	_, err = database.DB.Exec(execVerifyCode, verifyCode, AuthUserData.ID)
	if err != nil {
		log.Printf("Ошибка при добавлении верификационного кода в базу данных: %v", err)
		return err
	}
	log.Println("Код успешно добавлен")

	// Дополнительный код для регистрации пользователя, если он не существует
	// Например, хеширование пароля и создание нового пользователя в базе данных

	return nil
}
