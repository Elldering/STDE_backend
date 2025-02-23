package validation

import (
	"errors"
	"fmt"
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
	re := regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)(?=.*[!@#$%^&*()\-_=+{};:,<.>])[A-Za-z\d!@#$%^&*()\-_=+{};:,<.>]{8,}$`)
	return re.MatchString(password)
}

// CheckEmailOrPhoneNumber проверяет, является ли входная строка email или номером телефона.
// Возвращает "email", "phone" или ошибку, если строка не соответствует ни одному из форматов.
func CheckEmailOrPhoneNumber(login string) (string, error) {
	if ValidateEmail(login) {
		return "email", nil
	}
	if ValidatePhoneNumber(login) {
		return "phone", nil
	}
	return "", errors.New("некорректный логин: строка должна быть email или номером телефона")
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
