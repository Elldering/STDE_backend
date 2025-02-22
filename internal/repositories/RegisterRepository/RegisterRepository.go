package RegisterRepository

import (
	"STDE_proj/internal/models"
	"STDE_proj/utils/database"
	"STDE_proj/utils/smtp_sender"
	"STDE_proj/utils/validation"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func Register(login string, password string) error {
	if database.DB == nil {
		log.Println("Ошибка: подключение к базе данных не инициализировано")
		return fmt.Errorf("подключение к базе данных не инициализировано")
	}

	var data models.AuthUser
	var email, phoneNumber sql.NullString
	query := database.DB.QueryRow("SELECT id, email, phone_number FROM auth_user WHERE email = $1 OR phone_number = $2", login, login)
	err := query.Scan(&data.ID, &email, &phoneNumber)

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

	// Проверка, является ли username email или номером телефона

	// Проверяем, является ли login email
	isValidEmail, err := validation.ValidateEmail(login)
	if err != nil {
		return fmt.Errorf("ошибка при проверке email: %v", err)
	}
	if isValidEmail {
		data.Email = login
		return nil
	}

	// Проверяем, является ли login номером телефона
	isValidPhone, err := validation.ValidatePhoneNumber(login)
	if err != nil {
		return fmt.Errorf("ошибка при проверке номера телефона: %v", err)
	}
	if isValidPhone {
		data.PhoneNumber = login
		return nil
	}

	// Проверяем, не пустые ли поля login и password
	isValidFields, err := validation.ValidEmptyField(login, password)
	if err != nil {
		return fmt.Errorf("ошибка при проверке полей: %v", err)
	}
	if isValidFields {
		// Обработка случая, когда поля не пустые
		return nil
	}

	// Если ни одно из условий не выполнено, возвращаем ошибку
	return fmt.Errorf("некорректный email или номер телефона")

	// SQL-запрос для вставки данных в таблицу auth_user
	execAuthUser := "INSERT INTO auth_user (email, password) VALUES ($1, $2) RETURNING id"
	err = database.DB.QueryRow(execAuthUser, data.Email, password).Scan(&data.ID)
	if err != nil {
		log.Printf("Ошибка при добавлении пользователя в базу данных: %v", err)
		return err
	}

	log.Println("Пользователь успешно зарегистрирован")

	verifyCode := smtp_sender.GenerateCode()

	err = smtp_sender.SendEmail(login, verifyCode)
	if err != nil {
		log.Fatalf("Ошибка при отправке письма: %v", err)
	}

	// SQL-запрос для вставки данных в таблицу verify_code
	execVerifyCode := "INSERT INTO verify_code (code, auth_user_id) VALUES ($1, $2)"
	_, err = database.DB.Exec(execVerifyCode, verifyCode, data.ID)
	if err != nil {
		log.Printf("Ошибка при добавлении верификационного кода в базу данных: %v", err)
		return err
	}
	log.Println("Код успешно добавлен")

	// Дополнительный код для регистрации пользователя, если он не существует
	// Например, хеширование пароля и создание нового пользователя в базе данных

	return nil
}
