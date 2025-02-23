package validation

import (
	"errors"
	"fmt"
	"github.com/dlclark/regexp2"
	"regexp"
)

// ValidateEmail проверяет, является ли email корректным
func ValidateEmail(email string) bool {
	if len(email) < 0 {
		return false
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// ValidatePhoneNumber проверяет, является ли номер телефона корректным
func ValidatePhoneNumber(phone string) bool {
	if len(phone) < 0 {
		return false
	}
	re := regexp.MustCompile(`^\+?[1-9]\d{10,14}$`)
	return re.MatchString(phone)
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	re := regexp2.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)(?=.*[!@#$%^&*()\-_=+{};:,<.>])[A-Za-z\d!@#$%^&*()\-_=+{};:,<.>]{8,}$`, 0)
	match, _ := re.MatchString(password)
	return match
}

func CheckEmailOrPhoneNumber(login string) (string, error) {
	if ValidateEmail(login) {
		return "email", nil
	}
	if ValidatePhoneNumber(login) {
		return "phone", nil
	}
	return "", errors.New("некорректный логин")
}

func ValidateEmptyField(fields ...string) (bool, error) {
	for _, field := range fields {
		if field == "" {
			return false, fmt.Errorf("поле не может быть пустым")
		}
	}
	return true, nil
}
