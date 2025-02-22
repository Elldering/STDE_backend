package validation

import (
	"errors"
	"fmt"
	"regexp"
)

// ValidateEmail проверяет, является ли email корректным

import (
	"net/mail"
	"strings"
)

func ValidateEmail(email string) (bool, error) {
	// Проверка длины
	if len(email) < 5 || len(email) > 254 {
		return false, errors.New("длина email должна быть от 5 до 254 символов")
	}

	// Проверка с использованием регулярного выражения
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return false, errors.New("некорректный формат email")
	}

	// Проверка доменной части
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false, errors.New("email должен содержать символ @")
	}
	domain := parts[1]
	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
		return false, errors.New("доменная часть не может начинаться или заканчиваться на точку")
	}
	if !strings.Contains(domain, ".") {
		return false, errors.New("доменная часть должна содержать хотя бы одну точку")
	}

	// Проверка с использованием net/mail
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, fmt.Errorf("некорректный email: %v", err)
	}

	return true, nil
}

// ValidatePhoneNumber проверяет, является ли номер телефона корректным
func ValidatePhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+?[1-9]\d{10,14}$`)
	return re.MatchString(phone)
}

func ValidEmptyField(fields ...string) (bool, error) {
	for _, field := range fields {
		if field == "" {
			return false, fmt.Errorf("поле не может быть пустым")
		}
	}
	return true, nil
}
