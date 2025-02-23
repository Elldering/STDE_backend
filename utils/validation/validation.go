package validation

import (
	"STDE_proj/internal/models"
	"fmt"
	"github.com/dlclark/regexp2"
	"regexp"
)

// ValidateEmail проверяет, является ли email корректным.
func ValidateEmail(email string) bool {
	if len(email) == 0 {
		return false
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// ValidatePhoneNumber проверяет, является ли номер телефона корректным.
func ValidatePhoneNumber(phone string) bool {
	if len(phone) == 0 {
		return false
	}
	re := regexp.MustCompile(`^\+?[1-9]\d{10,14}$`)
	return re.MatchString(phone)
}

// ValidatePassword проверяет, является ли пароль корректным.
// Пароль должен содержать минимум 8 символов, включая буквы, цифры и специальные символы.
func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	re := regexp2.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)(?=.*[!@#$%^&*()\-_=+{};:,<.>])[A-Za-z\d!@#$%^&*()\-_=+{};:,<.>]{8,}$`, 0)
	match, _ := re.MatchString(password)
	return match
}

// CheckEmailOrPhoneNumber проверяет, является ли входная строка email или номером телефона.
// Возвращает "email", "phone" или ошибку, если строка не соответствует ни одному из форматов.
func CheckEmailOrPhoneNumber(data *models.AuthUser) error {
	if ValidateEmail(data.Login) {
		data.TypeLogin = "email"
		return nil
	}
	if ValidatePhoneNumber(data.Login) {
		data.TypeLogin = "phone_number"
		return nil
	}
	return fmt.Errorf("логин должен быть email или номером телефона")
}

// ValidateEmptyFields проверяет, что ни одно из переданных полей не пустое.
// Возвращает true, если все поля заполнены, или ошибку, если хотя бы одно поле пустое.
func ValidateEmptyFields(fields ...string) (bool, error) {
	for _, field := range fields {
		if field == "" {
			return false, fmt.Errorf("поле не может быть пустым")
		}
	}
	return true, nil
}
