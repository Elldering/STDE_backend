package validation

import (
	"regexp"
)

// ValidateEmail проверяет, является ли email корректным
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// ValidatePhoneNumber проверяет, является ли номер телефона корректным
func ValidatePhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+?[1-9]\d{10,14}$`)
	return re.MatchString(phone)
}
